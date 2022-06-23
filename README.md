# ethermint-gateway：

## Abstract
ethermint-gateway is the gateway service of ethermint blockchain. Its purpose is to better help ordinary programmers 
who do not understand the blockchain to complete the on-chain operation through etherming-gateway.

## Project launch:

ethermint-gateway completes compilation or runtime with go build/run. The default service listening address is 127.0.0.1:8000. To change the value, run the -addr command. 
go run ethermint-gateway -addr 127.0.0.1:8080 


## configuration file:
We put the ethermint-gateway configuration file in config.yaml


## API:

### user group

- 1./v1/user/addAccount :Add account to the database

      method post
      request:
        {
        "address":"0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282",
        "name":"mykey1"
        }

- 2./v1/user/getAccount : Query account by user_name

      method get
      name="mykey1" or address="0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282"


- 3./v1/contract/add : Storage contract account

      method post
      request:
        {
          "address":"0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282",
          "name":"myContract",
          "abi":"this is a invalid abi"
        }
        
- 4./v1/contract/get : Get contract info by contract name or contract address

      method get
      name="mykey1" or address="0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282"
      
      
      
### ethermint base group

- 1./v1/eth/base/getBalance: Query the user's non-contract account balance on ethermint blockchain by user name.

      method get
      name="mykey1" and block_number="" //block_number is used to query the account balance at a certain height, default ""
     
- 2./v1/eth/base/getBlockByNumber: Query block information by block height

      method get
      block_height="1000" //block_height is the height of a block, and we use string here. default ""  
      
- 3./v1/eth/base/getBlockByHash: Query block information using block hash

      method get
      block_hash="0x6b321c9497642b19247e5e1a300d32ffb3dee9dacbfc45e2df1e89e515c6fc95" //block_hash refers to the hash of the block we want to query, in this case we use string.
      
- 4./v1/eth/base/getBlockNumber: Query the last block of the current blockchain network

      method get

- 5./v1/eth/base/getTransactionReceipt: Query the transaction result by trading address

      method get
      tx_id="0x8e114fe45af750cb4cedfc9284168d8cd6c71506dbe2b5310e2bc541fed7651a"
      
- 6./v1/eth/base/getTransactionByHash: Query transaction information by trading address

      method get
      tx_id="0x8e114fe45af750cb4cedfc9284168d8cd6c71506dbe2b5310e2bc541fed7651a"

- 7./v1/eth/base/getNonceAt: Get the next nonce value of the account

      method get
      name="mykey1" and block_number="" //block_numberis used to query the next nonce value of a certain height, default ""

- 8./v1/eth/base/getCode: Get contract code

      method get
      name="myContract" //contract name
      block_number="" //block_number is used to query the contract code of a certain height, default ""
      
- 9./v1/eth/base/getEstimateGas: Get estimated Gas to execute the contract
 
       method post
       request:
         {
           "from":"mykey",
           "to":"testContract",
           "message":"setGreeting",
           "parameters":["nihao, for postman"]
         }
         
- 10./v1/eth/base/call : Execute the simulation contract and get the simulation result
 
       method post
       request:
         {
           "from":"mykey",
           "to":"testContract",
           "message":"setGreeting",
           "parameters":["nihao, for postman"]
         }
         
- 11./v1/eth/base/prepare : Hash the contract transaction and return the hash, which can be signed by the front-end user. 
 
       method post
       request:
         {
           "from":"mykey",
           "to":"mykey1",
           "value":4990324
         }
         
- 12./v1/eth/base/transfer : Send contract transaction with added signature
 
       method post
       request:
         {
           "from":"mykey",
           "to":"mykey1",
           "value":4990324,
           "sign":"f233ecafceca4e464d38cf0f91dcdbb50092fca2101d6abc5a9e09b026d331b8782c59b32de6741320751244d5dce0adfe2bfac526be17a1f8fd6ab7cedd721a00"
         }
         
- 13./v1/eth/base/sendRawTransaction : Send an encoded transaction message.  The transaction is a signed and encoded transaction.
 
       method post
       request:
         {
           "tx":"0xf8c6820b0d64830f4240943bfc61cba24a835f99a5855e613b239c7110b61980b864a4136862000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000136e6968616f2c20666f7220706f73746d616e3300000000000000000000000000279f31a587e104167bdf39ddfdb9b27006a75ccc8cff12d5713dd3e1d59030e2d0a0437b9c62164ee9d7c3daf88b66677153e957cfa296ee36fbeecde6270b1aecc0"
         }

         
### ethermint contract group

- 14./v1/eth/contract/view : Execute the simulation contract and get the simulation result.
 
       method post
       request:
         {
           "from":"mykey",
           "to":"testContract",
           "message":"getGreeting",
           "parameters":["nihao, for postman"]
         }
         
- 15./v1/eth/contract/prepare : Hash the contract transaction and return the hash, which can be signed by the front-end user. 
 
       method post
       request:
         {
           "from":"mykey",
           "to":"testContract",
           "message":"setGreeting",
           "parameters":["nihao, for postman"]
         }
         
- 16./v1/eth/contract/do : Send contract transaction with added signature
 
       method post
       request:
         {
           "from":"mykey",
           "to":"testContract",
           "message":"setGreeting",
           "parameters":["nihao, for postman"],
           "sign":"ae754b76ab2039ce5e6372dba37102125dce95e490a9f980087cf5587f036b461989b4a99174908f43658edc363b3a04f2ce1cc9975c93f5ae733e61eebdf91900"
         }
         
         
 ### ethermint erc20 contract group    
 - 17./v1/eth/contract/erc20/name: Query the internal name of the contract by user name and contract name defined by us.
 
       method get
       name="mykey1"//call user  name and contract_name="" //contract_name contract name can not be empty
       
 - 18./v1/eth/contract/erc20/symbol: Query the internal symbol of the contract by user name and the contract name defined by us.
 
       method get
       name="mykey1"//call user name and contract_name="" //contract_name contract name can not be empty
       
 - 19./v1/eth/contract/erc20/decimals: Query the internal decimals of the contract by user name and the contract name defined by us.
  
        method get
        name="mykey1"//call user name and contract_name="" //contract_name contract name can not be empty
        
 - 20./v1/eth/contract/erc20/totalSupply: Query the internal totalSupply of the contract by user name and the contract name defined by us.
   
         method get
         name="mykey1"//call user name and contract_name="" //contract_name contract name can not be empty

 - 21./v1/eth/contract/erc20/balanceOf: Query account balance in ERC20 by user name and the contract name defined by us.
  
        method get
        name="mykey1"//call user name 
        contract_name="" //contract_name contract name can not be empty
        owner="0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282" //Account address to be queried internally in ERC20
        
        
 - 22./v1/eth/contract/erc20/allowance: Query the internal allowance of the contract by user name and the contract name defined by us.
   
         method get
         name="mykey1"//call user name 
         contract_name="" //contract_name contract name can not be empty
         owner="0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282" //Account address to be queried internally in ERC20
         spender="0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282" //Account address to be queried internally in ERC20
         
 - 23./v1/eth/contract/erc20/transfer/prepare : Hash the contract transaction and return the hash, which can be signed by the front-end user. 
  
        method post
        request:
          {
            "name":"mykey",
            "contract_name":"teteher",
            "to":"testContract",
            "value":10000
          }
          
 - 24./v1/eth/contract/erc20/transfer/do : Send contract transaction with added signature
  
        method post
        request:
          {
            "name":"mykey",
            "contract_name":"teteher",
            "to":"0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282",
            "value":10000,
            "sign":"40fb0946a4a96b04332e7c27ab213285a44f0a080b9ed9b1b2dead1bc4030f90034e94d13c972dd987debd4dcb9954859147e43ca12ccaf1e62e8235ecdb69c700"
          }
 - 25./v1/eth/contract/erc20/approve/prepare : Hash the contract transaction and return the hash, which can be signed by the front-end user. 
   
         method post
         request:
           {
             "name":"mykey",
             "contract_name":"teteher",
             "spender":"0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282",
             "value":10000
           }
           
 - 26./v1/eth/contract/erc20/approve/do : Send contract transaction with added signature
   
         method post
         request:
           {
             "name":"mykey",
             "contract_name":"teteher",
             "spender":"0x533409D2BE9D874A57D4DC7EAD41D399FB2BE282",
             "value":10000,
             "sign":"20246b2089ce9ddfb3ecc62d81b0855e0ab570591b5c5e32d976dd21c68fcbde62644299616bce7170012b75d98a3871cfe1c17a7e986cb0db5e4e441c523a7a01"
           }
  
 - 27./v1/eth/contract/erc20/transferFrom/prepare : Hash the contract transaction and return the hash, which can be signed by the front-end user. 
     
           method post
           request:
             {
               "name":"mykey",
               "contract_name":"teteher",
               "from":"0x9321A2F1DC2B417B7165870A27873D2A7813A616",
               "to":"0xC9D057F7EE6EEAF68A52158B1EFE126D91923271",
               "value":1000
             }
             
 - 28./v1/eth/contract/erc20/transferFrom/do : Send contract transaction with added signature
     
           method post
           request:
             {
               "name":"mykey",
               "contract_name":"teteher",
               "from":"0x9321A2F1DC2B417B7165870A27873D2A7813A616",
               "to":"0xC9D057F7EE6EEAF68A52158B1EFE126D91923271",
               "value":1000,
               "sign":"67efbdbe703e1728214b44bc470886362fa6cb79f461633fb944739b6eb8e3826ac2e19fcc8b2cf31a7c5a5ce474268df2f18441bb46e186c1155288afb4f68601"
             }   
