xls2csv $1 | sed -e's/"//g'| sed '1d' |
while IFS=, read aadharid bioinfo income
do
  echo "$bioinfo">>temp.txt;
  mkdir $aadharid;
  accountid=$(sh newaccount.sh $aadharid $2 temp.txt);
  prefix='Address: {';
  suffix='}';
  accountid=$(echo "$accountid" | sed -e "s/^$prefix//" -e "s/$suffix$//");
  userdetails=$(echo $accountid $aadharid);
  echo $userdetails >> userAccountInfo.txt;
  sh registerCitizens.sh $accountid $bioinfo $income;
  rm temp.txt;
done
