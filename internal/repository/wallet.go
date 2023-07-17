package repository

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"net/mail"
	"ownify_api/internal/domain"
	"ownify_api/internal/dto"
	"ownify_api/internal/utils"
	"strings"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/mnemonic"
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"
)

type WalletQuery interface {
	AddNewAccount(role string, userId string, email string) (*string, error)
	RegisterNewAccount(walletAddress string, userId string) (*string, error)
	GetMyAccounts(email string, userId string) ([]string, error)
	MintOwnify(email string, pubKey string, products []dto.BriefProduct, net string) ([]uint64, error)
	UpdatePinCode(role string, email string, newPinCode string) error

	MakeTx(rawTx []byte, net string) (*string, *uint64, error)

	SendOwnify(email string, assetIds []uint64, sender string, receiver string, net string) (*string, error)

	DeleteOwnify(email string, assetIds []uint64, owner string, net string) (*string, error)
}

type walletQuery struct {
	ProductQueryService ProductQuery
}

func (w *walletQuery) AddNewAccount(
	role string,
	userId string,
	email string,
) (*string, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}
	tableName := BusinessTableName
	if role == domain.PersonalWallet {
		tableName = UserTableName
	}
	//get user pin code hash from wallets table
	var pin string
	sqlBuilder := utils.NewSqlBuilder()
	sql, err := sqlBuilder.Select(tableName, []string{"pin"}, []utils.Tuple{{Key: "email", Val: email}}, "=", "OR")
	if err != nil {
		return nil, err
	}

	err = DB.QueryRow(*sql).Scan(&pin)
	if err != nil {
		return nil, err
	}

	//create new EOA in algorand.
	newAcc := crypto.GenerateAccount()
	mnemonic, err := mnemonic.FromPrivateKey(newAcc.PrivateKey)
	if err != nil {
		return nil, err
	}
	pubKey := newAcc.Address.String()

	//encrypt mnemonic.
	cipher, err := utils.Encrypt(mnemonic, pin)
	if err != nil {
		return nil, err
	}

	//inset to wallet table.
	cols := []string{"chain_id", "user_id", "pub_addr", "email", "user_role", "seed_cipher"}
	values := []interface{}{0, userId, pubKey, email, role, cipher}

	query, err := sqlBuilder.Insert(WalletTableName, cols, values)
	if err != nil {
		return nil, err
	}
	_, err = DB.Exec(*query)
	if err != nil {
		return nil, err
	}
	return &pubKey, nil
}

func (w *walletQuery) GetMyAccounts(
	email string,
	userId string,
) ([]string, error) {
	var accounts = []string{}
	sqlBuilder := utils.NewSqlBuilder()
	sql, err := sqlBuilder.Select(WalletTableName, []string{
		"pub_addr",
	}, []utils.Tuple{{Key: "email", Val: email}, {Key: "user_id", Val: userId}}, "=", "OR")
	if err != nil {
		return []string{}, err
	}

	rows, err := DB.Query(*sql)
	for rows.Next() {
		var acc string
		err = rows.Scan(&acc)
		if err == nil {
			accounts = append(accounts, acc)
		}
	}
	if err != nil {
		return []string{}, err
	}
	return accounts, nil
}

func (w *walletQuery) MintOwnify(
	email string,
	pubKey string,
	products []dto.BriefProduct,
	net string,
) ([]uint64, error) {
	prv, err := recoverAccProcess(email, pubKey)
	if err != nil {
		return nil, err
	}

	//algorand client initialize
	client, _, err := NewClient(net)
	if err != nil {
		return nil, err
	}

	// Split the transactions into groups of 15.
	chunkSize := 15
	var chunks [][]dto.BriefProduct
	for i := 0; i < len(products); i += chunkSize {
		end := i + chunkSize

		if end > len(products) {
			end = len(products)
		}

		chunks = append(chunks, products[i:end])
	}

	var assetIndices []uint64
	results := make(chan []uint64, len(chunks)) // Channel to collect results
	errors := make(chan error, len(chunks))     // Channel to collect errors

	for _, chunk := range chunks {
		go func(chunk []dto.BriefProduct) { // Launch a goroutine for each chunk
			var assetChunkIndices []uint64
			txns := []types.Transaction{}
			for _, product := range chunk {
				note, err := json.Marshal(product)
				if err != nil {
					errors <- err
					return
				}
				txParams, err := client.SuggestedParams().Do(context.Background())
				if err != nil {
					errors <- err
					return
				}
				metaHash := utils.Hash(fmt.Sprintf("%v", note))
				txn, err := transaction.MakeAssetCreateTxn(pubKey,
					note,
					txParams, 1, 0,
					false, pubKey, pubKey, pubKey, pubKey,
					domain.OwnifyAssetName, domain.OwnifyAssetUnit, domain.OwnifyAssetMetaUrl, metaHash)

				if err != nil {
					errors <- err
					return
				}
				txns = append(txns, txn)
			}

			groupedTxs, err := transaction.AssignGroupID(txns, pubKey)
			if err != nil {
				errors <- err
				return
			}
			var stxs []byte
			for _, txn := range groupedTxs {
				_, stx, _ := crypto.SignTransaction(prv, txn)
				stxs = append(stxs, stx...)
			}

			pendingTxID, err := client.SendRawTransaction(stxs).Do(context.Background())

			if err != nil {
				errors <- err
				return
			}
			confirmedTx, err := transaction.WaitForConfirmation(client, pendingTxID, 4, context.Background())
			if err != nil {
				errors <- err
				return
			}
			endIndex := confirmedTx.AssetIndex + uint64(len(chunk))

			// add to database
			for i := 0; i < len(chunk); i++ {
				chunk[i].AssetId = int64(confirmedTx.AssetIndex + uint64(i))
				chunk[i].Owner = pubKey
			}
			err = w.ProductQueryService.AddProducts(chunk, net, false)
			if err != nil {
				errors <- err
				return
			}
			// Add the asset indices of the current chunk to the final result.
			assetChunkIndices = append(assetChunkIndices, utils.MakeRange(confirmedTx.AssetIndex, endIndex)...)
			results <- assetChunkIndices
		}(chunk)
	}

	// Collect results and errors
	for i := 0; i < len(chunks); i++ {
		select {
		case res := <-results:
			assetIndices = append(assetIndices, res...)
		case err := <-errors:
			return nil, err
		}
	}

	return assetIndices, nil
}

func processChunks(assetIds []uint64, chunkSize int, processFunc func([]uint64) (*string, error)) (*string, error) {
	var chunks [][]uint64
	for i := 0; i < len(assetIds); i += chunkSize {
		end := i + chunkSize
		if end > len(assetIds) {
			end = len(assetIds)
		}

		chunks = append(chunks, assetIds[i:end])
	}

	var allTxns []string

	for _, chunk := range chunks {
		txIds, err := processFunc(chunk)
		if err != nil {
			return nil, err
		}
		allTxns = append(allTxns, *txIds)
	}

	totalTxId := strings.Join(allTxns, ",")
	return &totalTxId, nil
}

func (w *walletQuery) SendOwnify(
	email string,
	assetIds []uint64,
	sender string,
	receiver string,
	net string,
) (*string, error) {
	prv, err := recoverAccProcess(email, sender)
	if err != nil {
		return nil, err
	}

	//algorand client initialize
	client, _, err := NewClient(net)
	if err != nil {
		return nil, err
	}

	sendChunk := func(chunk []uint64) (*string, error) {
		txns := []types.Transaction{}
		for _, assetId := range chunk {
			txParams, err := client.SuggestedParams().Do(context.Background())
			if err != nil {
				return nil, err
			}
			txn, err := transaction.MakeAssetTransferTxn(sender,
				receiver, 1,
				[]byte(""),
				txParams,
				"",
				assetId,
			)

			if err != nil {
				return nil, err
			}
			txns = append(txns, txn)
		}

		return sendGroupedTransactions(client, prv, sender, txns)
	}
	return processChunks(assetIds, 15, sendChunk)
}

func (w *walletQuery) DeleteOwnify(
	email string,
	assetIds []uint64,
	owner string,
	net string,
) (*string, error) {
	prv, err := recoverAccProcess(email, owner)
	if err != nil {
		return nil, err
	}

	//algorand client initialize
	client, _, err := NewClient(net)
	if err != nil {
		return nil, err
	}

	// Split the assetIds into chunks of 15.
	chunkSize := 15
	var chunks [][]uint64
	for i := 0; i < len(assetIds); i += chunkSize {
		end := i + chunkSize

		if end > len(assetIds) {
			end = len(assetIds)
		}

		chunks = append(chunks, assetIds[i:end])
	}

	results := make(chan *string, len(chunks)) // Channel to collect results
	errors := make(chan error, len(chunks))    // Channel to collect errors

	for _, chunk := range chunks {
		go func(chunk []uint64) { // Launch a goroutine for each chunk
			txns := []types.Transaction{}

			for _, assetId := range chunk {

				txParams, err := client.SuggestedParams().Do(context.Background())
				if err != nil {
					errors <- err
					return
				}
				txn, err := transaction.MakeAssetDestroyTxn(owner,
					[]byte(""),
					txParams,
					assetId,
				)

				if err != nil {
					errors <- err
					return
				}
				txns = append(txns, txn)
			}
			txId, err := sendGroupedTransactions(client, prv, owner, txns)
			if err != nil {
				errors <- err
				return
			}
			// delete operation
			w.ProductQueryService.DeleteProducts(chunk, net)
			results <- txId
		}(chunk)
	}

	// Collect results and errors
	var lastTxId *string
	for i := 0; i < len(chunks); i++ {
		select {
		case res := <-results:
			lastTxId = res
		case err := <-errors:
			return nil, err
		}
	}

	return lastTxId, nil
}

func sendGroupedTransactions(client *algod.Client, prv ed25519.PrivateKey, sender string, txns []types.Transaction) (*string, error) {
	groupedTxs, err := transaction.AssignGroupID(txns, sender)
	if err != nil {
		return nil, err
	}
	var stxs []byte
	for _, txn := range groupedTxs {
		_, stx, _ := crypto.SignTransaction(prv, txn)
		stxs = append(stxs, stx...)
	}

	if err != nil {
		return nil, err
	}

	pendingTxID, err := client.SendRawTransaction(stxs).Do(context.Background())

	if err != nil {
		return nil, err
	}
	_, err = transaction.WaitForConfirmation(client, pendingTxID, 4, context.Background())
	if err != nil {
		return nil, err
	}
	return &pendingTxID, nil
}

func recoverAccProcess(email, pubKey string) (ed25519.PrivateKey, error) {
	cipherR := make(chan domain.Result[string])
	pinR := make(chan domain.Result[string])
	defer close(pinR)
	defer close(cipherR)
	go func() {
		var cipher string
		sqlBuilder := utils.NewSqlBuilder()
		sql, err := sqlBuilder.Select(WalletTableName, []string{
			"seed_cipher",
		}, []utils.Tuple{{Key: "email", Val: email}, {Key: "pub_addr", Val: pubKey}}, "=", "AND")
		if err != nil {
			cipherR <- domain.Result[string]{Err: err}
			return
		}
		err = DB.QueryRow(*sql).Scan(&cipher)
		if err != nil {
			cipherR <- domain.Result[string]{Err: err}
			return
		}
		seed := ""
		err = DB.QueryRow(*sql).Scan(&seed)
		if seed == "" {
			cipherR <- domain.Result[string]{Err: err}
			return
		}
		cipherR <- domain.Result[string]{Val: seed}
	}()

	//get pin code from business table.
	go func() {
		pin := ""
		sqlBuilder := utils.NewSqlBuilder()
		sql, err := sqlBuilder.Select(BusinessTableName, []string{
			"pin",
		}, []utils.Tuple{{Key: "email", Val: email}}, "=", "AND")
		if err != nil {
			pinR <- domain.Result[string]{Err: err}
			return
		}
		err = DB.QueryRow(*sql).Scan(&pin)
		if err != nil {
			pinR <- domain.Result[string]{Err: err}
			return
		}
		pinR <- domain.Result[string]{Val: pin}
	}()

	pin := <-pinR
	cipher := <-cipherR

	if !pin.Ok() {
		return nil, pin.Err
	}
	if !cipher.Ok() {
		return nil, pin.Err
	}

	// decrypt cipher for recover account.
	seed, err := utils.Decrypt(cipher.Val, pin.Val)
	if err != nil {
		return nil, pin.Err
	}
	prv, err := mnemonic.ToPrivateKey(seed)

	if err != nil {
		return nil, err
	}

	acc, err := crypto.AccountFromPrivateKey(prv)
	if acc.Address.String() != pubKey {
		return nil, fmt.Errorf("Address doesn't match")
	}
	return prv, err
}

func (w *walletQuery) UpdatePinCode(role string, userId string, newPinCode string) error {
	//var cipher string
	///var email string
	// err := pgQb().Select("pub_addr", "email").Where(sq.Eq{"pub_addr": pubKey}).From("ownify.wallets").QueryRow().Scan(&email, &cipher)
	// if err != nil {
	// 	return []string{}, err
	// }

	return nil
}

func (w *walletQuery) MakeTx(rawTx []byte, net string) (*string, *uint64, error) {
	//algorand client initialize
	client, _, err := NewClient(net)
	pendingTxID, err := client.SendRawTransaction(rawTx).Do(context.Background())

	confirmedTx, err := transaction.WaitForConfirmation(client, pendingTxID, 4, context.Background())
	if err != nil {
		return nil, nil, err
	}
	endIndex := confirmedTx.AssetIndex
	return &pendingTxID, &endIndex, nil
}

func (w *walletQuery) RegisterNewAccount(walletAddress string, userId string) (*string, error) {
	cols := []string{"chain_id", "pub_addr", "user_id"}
	values := []interface{}{1, walletAddress, userId}

	sqlBuilder := utils.NewSqlBuilder()
	query, err := sqlBuilder.Insert("wallets", cols, values)
	if err != nil {
		return nil, err
	}
	_, err = DB.Exec(*query)
	if err != nil {
		return nil, err
	}
	return &walletAddress, nil
}
