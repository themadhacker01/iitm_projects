var abi = eth.contract([{"constant":true,"inputs":[],"name":"curdealer","outputs":[{"name":"subsidyCollected","type":"uint16"},{"name":"remainingGoods","type":"uint16"},{"name":"outstandingGoods","type":"uint16"},{"name":"bioInfo","type":"uint16"},{"name":"stateOfficial","type":"address"},{"name":"dealerPermission","type":"uint8"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"initial_balance","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"printDealerDetails","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"rate","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"statedb","outputs":[{"name":"subsidyCollected","type":"uint16"},{"name":"bioInfo","type":"uint16"},{"name":"statePermission","type":"uint8"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"curstateofficial","outputs":[{"name":"subsidyCollected","type":"uint16"},{"name":"bioInfo","type":"uint16"},{"name":"statePermission","type":"uint8"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_amount","type":"uint16"},{"name":"_callerIndex","type":"uint16"},{"name":"_bioInfo","type":"uint16"}],"name":"payState","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"monthlySubsidyAlloc","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint16"}],"name":"stateIndex","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_callerIndex","type":"uint16"},{"name":"_amount","type":"uint16"},{"name":"_citizenIndex","type":"uint16"},{"name":"_bioInfo","type":"uint16"}],"name":"payDealer","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"center_admin","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"citizencount","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint16"}],"name":"dealerIndex","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint16"}],"name":"citizenIndex","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_citizenID","type":"address"}],"name":"allocateSubsidy","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"subsidy_budget","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"curcitizen","outputs":[{"name":"bioInfo","type":"uint16"},{"name":"bpl","type":"bool"},{"name":"subsidyBalance","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"final_balance","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"dealerdb","outputs":[{"name":"subsidyCollected","type":"uint16"},{"name":"remainingGoods","type":"uint16"},{"name":"outstandingGoods","type":"uint16"},{"name":"bioInfo","type":"uint16"},{"name":"stateOfficial","type":"address"},{"name":"dealerPermission","type":"uint8"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_accountAddress","type":"address"},{"name":"_aadharID","type":"string"},{"name":"_bioInfo","type":"uint16"}],"name":"registerStateOfficial","outputs":[{"name":"","type":"string"}],"payable":true,"type":"function"},{"constant":false,"inputs":[],"name":"printCustomerDetails","outputs":[{"name":"","type":"address"},{"name":"","type":"bool"},{"name":"","type":"uint16"},{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"stateofficialcount","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_accountAddress","type":"address"},{"name":"_bioInfo","type":"uint16"},{"name":"_income","type":"uint32"}],"name":"registerCitizen","outputs":[{"name":"","type":"string"}],"payable":true,"type":"function"},{"constant":false,"inputs":[{"name":"_callerIndex","type":"uint16"},{"name":"_accountAddress","type":"address"},{"name":"_aadharID","type":"string"},{"name":"_bioInfo","type":"uint16"}],"name":"registerDealer","outputs":[{"name":"","type":"string"}],"payable":true,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"citizendb","outputs":[{"name":"bioInfo","type":"uint16"},{"name":"bpl","type":"bool"},{"name":"subsidyBalance","type":"uint16"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_amount","type":"uint16"},{"name":"_stateIndex","type":"uint16"},{"name":"_bioInfo","type":"uint16"}],"name":"payCenter","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"dealercount","outputs":[{"name":"","type":"uint16"}],"payable":false,"type":"function"},{"inputs":[{"name":"_subsidy_budget","type":"uint16"},{"name":"_rate","type":"uint16"}],"payable":false,"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"p1","type":"address"},{"indexed":false,"name":"p2","type":"bool"},{"indexed":false,"name":"p3","type":"uint16"},{"indexed":false,"name":"p4","type":"uint16"}],"name":"printCustomer","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"p1","type":"address"},{"indexed":false,"name":"subsidyCollected","type":"uint16"},{"indexed":false,"name":"remainingGoods","type":"uint16"},{"indexed":false,"name":"outstandingGoods","type":"uint16"},{"indexed":false,"name":"bioInfo","type":"uint16"}],"name":"printDealer","type":"event"}]).at('0xcbaf6a8e610e1f7c90d282df2f657cc5d583ff48')
