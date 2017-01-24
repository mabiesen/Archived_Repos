//YWROBOT
//Compatible with the Arduino IDE 1.0
//Library version:1.1
#include <Wire.h> 
#include <LiquidCrystal_I2C.h>

LiquidCrystal_I2C lcd(0x27,20,4);  // set the LCD address to 0x27 for a 16 chars and 2 line display

void setup()
{
  lcd.init();                      // initialize the lcd 
  lcd.init();
  // Print a message to the LCD.

  

  pinMode(LED_BUILTIN, OUTPUT);
}


void loop()
{
  lcd.backlight();
  lcd.setCursor(3,0);
  lcd.print("Hello, world!");
  digitalWrite(LED_BUILTIN, HIGH);   // turn the LED on (HIGH is the voltage level)
  delay(1000);// wait for a second
  lcd.noBacklight();
  digitalWrite(LED_BUILTIN, LOW);    // turn the LED off by making the voltage LOW
  delay(1000);                       // wait for a second
}
