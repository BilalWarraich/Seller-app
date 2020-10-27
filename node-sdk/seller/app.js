'use strict';

//get libraries
const express = require('express');
var queue = require('express-queue');
const bodyParser = require('body-parser');
const request = require('request');
const path = require('path');

//create express web-app
const app = express();
const invoke = require('./invokeNetwork');
const query = require('./queryNetwork');
var _time = "T00:00:00Z";

//declare port
var port = process.env.PORT || 8000;
if (process.env.VCAP_APPLICATION) {
  port = process.env.PORT;
}

app.use(bodyParser.json({
  limit: '50mb', 
  extended: true

}));

app.use(bodyParser.urlencoded({
 limit: '50mb', 
 extended: true
}));

app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  next();
  });

//Using queue middleware
app.use(queue({ activeLimit: 30, queuedLimit: -1 }));

//run app on port
app.listen(port, function () {
  console.log('app running on port: %d', port);
});

//-------------------------------------------------------------
//----------------------  POST API'S    -----------------------
//-------------------------------------------------------------

app.post('/api/addSeller', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'addSeller',
    args: [

      req.body.sellerID,
      req.body.username,
      req.body.password

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Seller with ID: "+req.body.sellerID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addBuyer', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'addBuyer',
    args: [

      req.body.buyerID,
      req.body.username,
      req.body.password

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Buyer with ID: "+req.body.buyerID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addMarketer', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'addMarketer',
    args: [

      req.body.marketerID,
      req.body.username,
      req.body.password

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The marketer with ID: "+req.body.marketerID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addProduct', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'addProduct',
    args: [
      req.body.productID,
      req.body.sellerID,
      req.body.buyerID,
      req.body.productName,
      req.body.price,
      req.body.description,
      req.body.soldDate,
      req.body.status

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The product with ID: "+req.body.productID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateProduct', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'updateProduct',
    args: [

      req.body.productID,
      req.body.newStatus,
      req.body.soldDate,
      req.body.buyerID

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Product with ID: "+req.body.productID+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addMarketedProduct', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'addMarketedProduct',
    args: [


      req.body.mID,
      req.body.marketerID,
      req.body.productID,
      req.body.sellerID,
      req.body.soldDate,
      req.body.status

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Marketed Product with ID: "+req.body.mID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateMarketedProduct', async function (req, res) {

  var request = {
    chaincodeId: 'seller',
    fcn: 'updateMarketedProduct',
    args: [

      req.body.mID,
      req.body.newStatus,
      req.body.soldDate,
    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The MarketedProduct with ID: "+req.body.mID+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

//-------------------------------------------------------------
//----------------------  GET API'S  --------------------------
//-------------------------------------------------------------

app.get('/api/querySeller', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'querySeller',
    args: [
      req.query.username,
      req.query.password
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryBuyer', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryBuyer',
    args: [
      req.query.username,
      req.query.password
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketer', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketer',
    args: [
      req.query.username,
      req.query.password
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/querySellerByName', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'querySellerByName',
    args: [
      req.query.username
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/querySellerByID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'querySellerByID',
    args: [
      req.query.sellerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryAllSellers', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryAllSellers',
    args:[]
   
  };
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
    res.status(response.status).send(JSON.parse(response.message));
    else
    res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryBuyerByID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryBuyerByID',
    args: [
      req.query.buyerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketerByID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketerByID',
    args: [
      req.query.marketerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductByID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductByID',
    args: [
      req.query.productID
    
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProduct', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProduct',
    args: [
      req.query.status,
      req.query.productID
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProducts', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProducts',
    args: [
      req.query.sellerID,
      req.query.buyerID
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductbySellerID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductbySellerID',
    args: [
      req.query.sellerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductbyBuyerID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductbyBuyerID',
    args: [
      req.query.buyerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductbyProductName', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductbyProductName',
    args: [
      req.query.username
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductbyStatus', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductbyStatus',
    args: [
      req.query.status
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductbyDate', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductbyDate',
    args: [
      req.query.soldDate
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryProductbyDate', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryProductbyDate',
    args: [
      req.query.status,
      req.query.marketerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketedProductByID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketedProductByID',
    args: [
      req.query.mID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketedProductBySellerID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketedProductBySellerID',
    args: [
      req.query.sellerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketedProductByMarketerID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketedProductByMarketerID',
    args: [
      req.query.marketerID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketedProductByProductID', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketedProductByProductID',
    args: [
      req.query.productID
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketedProductByDate', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketedProductByDate',
    args: [
      req.query.soldDate
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryMarketedProductByStatus', async function (req, res) {

  const request = {
    chaincodeId: 'seller',
    fcn: 'queryMarketedProductByStatus',
    args: [
      req.query.status
        
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});