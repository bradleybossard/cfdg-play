files=`find output | grep cfdg`
for infile in $files; do
  outfile=img/`basename $infile .cfdg`.png
  cfdg $infile $outfile 
done;

files=`find img | grep png | sort`
backfiles=`find img | grep png | sort -r`
backforthfiles=`echo $files $backfiles`
gifname=`basename $(echo $files | cut -d" " -f1) .png`.gif
convert -delay 10 -loop 0 $backforthfiles img/$gifname
