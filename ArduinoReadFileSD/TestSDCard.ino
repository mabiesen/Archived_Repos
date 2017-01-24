#include <SD.h>
#include <SPI.h>

const int chipSelect = 10;
Sd2Card card;
SdFile root;
File myFile;

void setup()
{
  Serial.begin(9600);
  while (!Serial) {
    
  }
  
  if (!card.init(SPI_HALF_SPEED, chipSelect)) {
    Serial.println("initialization failed. Things to check:");
    Serial.println("* is a card inserted?");
    Serial.println("* is your wiring correct?");
    Serial.println("* did you change the chipSelect pin to match your shield or module?");
    return;
  } else {
    Serial.println("Wiring is correct and a card is present.");
  }

  if (!SD.begin(10)) {
    Serial.println("initialization failed!");
    return;
  }
  
  Serial.println("initialization done.");

  if (SD.exists("some.txt")){
    Serial.println("Some exists!");
    myFile = SD.open("some.txt");
    while (myFile.available()){
      Serial.write(myFile.read());
    }
    myFile.close();
  } else {
    Serial.println("Nothing exists");
  }
}   

void loop()
{}

