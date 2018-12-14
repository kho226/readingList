mkdir completed
cd completed
mkdir get
cd get
echo "GET http://localhost:8080/#/completed" | vegeta attack -workers=50 -duration=30s > test.bin
vegeta report test.bin
vegeta plot test.bin > test.html


cd ..
cd ..
mkdir active
cd active
mkdir get
cd get
echo "GET http://localhost:8080/#/active" | vegeta attack -workers=50 -duration=30s > test.bin
vegeta report test.bin
vegeta plot test.bin > test.html

cd ..
cd ..
mkdir all
cd all
mkdir get
cd get
echo "GET http://localhost:8080/#/all" | vegeta attack -workers=50 -duration=30s > test.bin
vegeta report test.bin
vegeta plot test.bin > test.html

cd ..
cd ..
mkdir readingList
cd readingList
mkdir get
cd get
echo "GET http://localhost:8080/readingList/" | vegeta attack -workers=50 -duration=30s > test.bin
vegeta report test.bin
vegeta plot test.bin > test.html

cd ..
cd ..
