pragma solidity ^0.4.7;

contract PDSAutomation{

      /*=======================================================================
        CONTRACT VARIABLES
        ========================================================================*/
      address public center_admin;
      uint16 public initial_balance;
      uint16 public final_balance;
      uint16 public subsidy_budget;
      uint16  public citizencount;
      uint16 public stateofficialcount;
      uint16 public dealercount;
      Citizen public curcitizen;
      StateOfficial public curstateofficial;
      Dealer public curdealer;
      uint16 public rate;

      struct Citizen
      {
          uint16 bioInfo;
          bool bpl;
          uint16 subsidyBalance;
      }

      struct Dealer
      {
          uint16 subsidyCollected;
          uint16 remainingGoods;
          uint16 outstandingGoods;
          uint16 bioInfo;
          address stateOfficial;
          uint8 dealerPermission;
      }

      struct StateOfficial
      {
          uint16 subsidyCollected;
          uint16 bioInfo;
          uint8 statePermission;
      }

      mapping (uint16 => address) public citizenIndex;
      mapping (uint16 => address) public stateIndex;
      mapping (uint16 => address) public dealerIndex;
      mapping (address => Citizen) public citizendb;
      mapping (address => Dealer) public dealerdb;
      mapping (address => StateOfficial) public statedb;

      event printCustomer(address p1, bool p2, uint16 p3, uint16 p4);
      event printDealer(address p1, uint16 subsidyCollected, uint16 remainingGoods, uint16 outstandingGoods, uint16 bioInfo);
      /*=======================================================================
        FUNCTION: CONSTRUCTOR TO INITALIZE CONTRACT
        ========================================================================*/
      function PDSAutomation(uint16 _subsidy_budget, uint16 _rate) public {

        center_admin = msg.sender;
        citizencount = 0;
        stateofficialcount = 0;
        dealercount = 0;
        subsidy_budget = _subsidy_budget;
        rate = _rate;
      }
      /*=======================================================================
        FUNCTION: PRINT CUSTOMER, DEALER, STATE, CENTER STATUS
        ========================================================================*/
      function printCustomerDetails() public returns(address, bool, uint16, uint16){

            //for (uint16 i=0; i<citizencount; i++)
            //{
                uint16 i=0;
                curcitizen = citizendb[citizenIndex[i]];
                return (citizenIndex[i], curcitizen.bpl, curcitizen.bioInfo, curcitizen.subsidyBalance);
            //}
      }
      function printDealerDetails() {

            for (uint16 i=0; i<dealercount; i++)
            {
                curdealer = dealerdb[dealerIndex[i]];
                printDealer(dealerIndex[i], curdealer.subsidyCollected, curdealer.remainingGoods, curdealer.outstandingGoods, curdealer.bioInfo);
            }
      }

     /*=======================================================================
        FUNCTION: Monthly Subsidy Allocator
        ========================================================================*/
      function monthlySubsidyAlloc() {

            initial_balance = subsidy_budget;
            for (uint16 i=0; i<citizencount; i++)
            {
               allocateSubsidy(citizenIndex[i]);
            }
      }
      /*=======================================================================
        FUNCTION: TO REGISTER A CITIZEN
        ========================================================================*/
      function registerCitizen(address _accountAddress, uint16 _bioInfo, uint32 _income) payable returns (string){
            if (msg.sender != center_admin)
            {
              throw;
            }
            citizendb[_accountAddress].bioInfo = _bioInfo;
            if (_income <= 100000)
              citizendb[_accountAddress].bpl = true;
            else
              citizendb[_accountAddress].bpl = false;

            citizenIndex[citizencount] = _accountAddress;
            citizencount++;
            return "registerCitizen";
      }
      /*=======================================================================
        FUNCTION: TO REGISTER A STATE OFFICIAL
        ========================================================================*/
      function registerStateOfficial(address _accountAddress, string _aadharID, uint16 _bioInfo) payable returns (string){
            if (msg.sender != center_admin)
            {
              throw;
            }
            statedb[_accountAddress].bioInfo = _bioInfo;
            statedb[_accountAddress].statePermission = 1;

            stateIndex[stateofficialcount] = _accountAddress;
            stateofficialcount++;
            return "registerStateOfficial";
      }
      /*=======================================================================
        FUNCTION: TO REGISTER A DEALER
        ========================================================================*/
      function registerDealer(uint16 _callerIndex, address _accountAddress, string _aadharID, uint16 _bioInfo) payable returns (string){

            if (statedb[stateIndex[_callerIndex]].statePermission != 1 || statedb[stateIndex[_callerIndex]].bioInfo != _bioInfo)
            {
              throw;
            }
            dealerdb[_accountAddress].bioInfo = _bioInfo;
            dealerdb[_accountAddress].dealerPermission = 1;
            dealerdb[_accountAddress].stateOfficial = stateIndex[_callerIndex];
            dealerdb[_accountAddress].remainingGoods = 200;

            dealerIndex[dealercount] = _accountAddress;
            dealercount++;
            return "registerDealer";
      }
     /*=======================================================================
        FUNCTION: TO ALLOCATE SUBSIDY
        ========================================================================*/
      function allocateSubsidy(address _citizenID) returns (string) {

          if (msg.sender != center_admin)
            throw;
          if (citizendb[_citizenID].bpl)
          {
            citizendb[_citizenID].subsidyBalance = 500;
            initial_balance -= 500;
          }
          else
          {
            citizendb[_citizenID].subsidyBalance = 250;
            initial_balance -= 250;
          }
          return "allocateSubsidy";
      }
      /*=======================================================================
        FUNCTION: TO PAY THE DEALER FROM <CITIZEN>s ACCOUNT
        ========================================================================*/
      function payDealer(uint16 _callerIndex, uint16 _amount, uint16 _citizenIndex, uint16 _bioInfo) returns (string){

            if (dealerdb[dealerIndex[_callerIndex]].dealerPermission != 1 || citizendb[citizenIndex[_citizenIndex]].bioInfo != _bioInfo)
            {
              throw;
            }
            if (citizendb[citizenIndex[_citizenIndex]].subsidyBalance < _amount * rate)
                throw;
            if (dealerdb[dealerIndex[_callerIndex]].remainingGoods < _amount)
                throw;
            dealerdb[dealerIndex[_callerIndex]].subsidyCollected = dealerdb[dealerIndex[_callerIndex]].subsidyCollected + _amount * rate;
            dealerdb[dealerIndex[_callerIndex]].remainingGoods = dealerdb[dealerIndex[_callerIndex]].remainingGoods - _amount;
            citizendb[citizenIndex[_citizenIndex]].subsidyBalance = citizendb[citizenIndex[_citizenIndex]].subsidyBalance - _amount * rate;
            return "payDealer";
      }
      /*=======================================================================
        FUNCTION: TO PAY THE STATE FROM <DEALER>s ACCOUNT
        ========================================================================*/
      function payState(uint16 _amount, uint16 _callerIndex, uint16 _bioInfo) returns (string){

            if (dealerdb[dealerIndex[_callerIndex]].dealerPermission != 1 || dealerdb[dealerIndex[_callerIndex]].bioInfo != _bioInfo)
            {
              throw;
            }
            if (dealerdb[dealerIndex[_callerIndex]].subsidyCollected < _amount * rate)
                throw;
            statedb[dealerdb[dealerIndex[_callerIndex]].stateOfficial].subsidyCollected = statedb[dealerdb[dealerIndex[_callerIndex]].stateOfficial].subsidyCollected + _amount * rate;
            dealerdb[dealerIndex[_callerIndex]].subsidyCollected = dealerdb[dealerIndex[_callerIndex]].subsidyCollected - _amount * rate;
            return "payState";
      }
      /*=======================================================================
        FUNCTION: TO PAY THE CENTER FROM <STATE>s ACCOUNT
        ========================================================================*/
      function payCenter(uint16 _amount, uint16 _stateIndex, uint16 _bioInfo) returns (string){

            if (statedb[stateIndex[_stateIndex]].statePermission != 1 || statedb[stateIndex[_stateIndex]].bioInfo != _bioInfo)
            {
              throw;
            }
            if (statedb[stateIndex[_stateIndex]].subsidyCollected < _amount)
                throw;
            statedb[stateIndex[_stateIndex]].subsidyCollected = statedb[stateIndex[_stateIndex]].subsidyCollected - _amount;
            final_balance = final_balance + _amount;
            return "payCenter";
      }
}
