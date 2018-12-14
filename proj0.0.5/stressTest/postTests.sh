mkdir readingList
cd readingList
mkdir post
cd post
echo "POST http://localhost:8080/readingList/" | vegeta attack -workers=50 -duration=30s > test.bin
vegeta report test.bin
vegeta plot test.bin > test.html

cd ..
cd ..