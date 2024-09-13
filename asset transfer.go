package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AssetContract struct {
    contractapi.Contract
}

type Asset struct {
    DEALERID    string `json:"dealerId"`
    MSISDN       string `json:"msisdn"`
    MPIN         string `json:"mpin"`
    BALANCE      int    `json:"balance"`
    STATUS       string `json:"status"`
    TRANSAMOUNT  int    `json:"transAmount"`
    TRANSTYPE    string `json:"transType"`
    REMARKS      string `json:"remarks"`
}

func (c *AssetContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerId string, msisdn string, mpin string, balance int, status string, transAmount int, transType string, remarks string) error {
    asset := Asset{
        DEALERID:   dealerId,
        MSISDN:     msisdn,
        MPIN:       mpin,
        BALANCE:    balance,
        STATUS:     status,
        TRANSAMOUNT: transAmount,
        TRANSTYPE:  transType,
        REMARKS:    remarks,
    }

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(dealerId, assetJSON)
}

func (c *AssetContract) QueryAsset(ctx contractapi.TransactionContextInterface, dealerId string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(dealerId)
    if err != nil {
        return nil, err
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", dealerId)
    }

    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

func main() {
    assetContract := new(AssetContract)
    chaincode, err := contractapi.NewChaincode(assetContract)
    if err != nil {
        fmt.Printf("Error creating asset-transfer-basic chaincode: %s", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting asset-transfer-basic chaincode: %s", err.Error())
    }
}
