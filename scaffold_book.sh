#chmod +x scaffold_book.sh
#./scaffold_book.sh

chapters=$(seq -w 1 15)
#chapters="chapter_{01..15}" 
#eval mkdir -p $chapters

for i in $chapters; do
  mkdir -p chapter_$i/sections/
  mkdir -p chapter_$i/examples/
done
