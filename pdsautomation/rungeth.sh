if [ ! -d "logs" ]; then
  mkdir logs
fi

if [ "$1" = "State" ]; then
  port="30304";
elif [ "$1" = "Center" ]; then
  port="30303";
else
  port="30305";
fi
  
geth --testnet --datadir ./$1 --networkid 10454 --mine --minerthreads 1 --unlock 0 --port $port console init CustomGensis.json 2> logs/trace.log
