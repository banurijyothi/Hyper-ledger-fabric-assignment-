const express = require('express');
const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

const app = express();
app.use(express.json());

const ccpPath = path.resolve(__dirname, 'connection.json');

async function queryAsset(dealerId) {
    const wallet = await Wallets.newFileSystemWallet('./wallet');
    const gateway = new Gateway();

    await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

    const network = await gateway.getNetwork('mychannel');
    const contract = network.getContract('asset-transfer-basic');

    const result = await contract.evaluateTransaction('QueryAsset', dealerId);
    await gateway.disconnect();

    return JSON.parse(result.toString());
}

app.get('/asset/:dealerId', async (req, res) => {
    try {
        const asset = await queryAsset(req.params.dealerId);
        res.json(asset);
    } catch (error) {
        res.status(500).send(error.message);
    }
});

app.listen(3000, () => {
    console.log('REST API server running on port 3000');
});
