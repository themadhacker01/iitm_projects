command=$(echo "'loadScript(\"./registerCitizens.js\"); abi.registerCitizen(0x$1, $2, $3,{from: eth.accounts[0], gas: 100000});'");
echo $command;
command=$(echo "geth --exec $command --port 30302 attach ipc:./Center/geth.ipc");
eval $command;
