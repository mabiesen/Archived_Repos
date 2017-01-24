#include "SD.h"
#include "TMRpcm.h"
#include "SPI.h"

const int chipSelect = 10;
TMRpcm tmrpcm;
Sd2Card card;
SdFile root;
File myFile;

void setup(){
  tmrpcm.speakerPin = 9;
  Serial.begin(9600);
  
  while (!Serial){
    
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

  
  if (!SD.begin(chipSelect)) {
    Serial.println("initialization failed!");
    return;
  }

  if (SD.exists("21.wav")){
    Serial.println("Santa exists!");
  }


tmrpcm.setVolume(6);
tmrpcm.play("21.wav");
}

void loop() {
  // put your main code here, to run repeatedly:

}
