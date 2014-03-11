files=`find output | grep cfdg`
for infile in $files; do
  outfile=img/`basename $infile .cfdg`.png
  cfdg $infile $outfile 
done;
