Directions online for the Funduino 1602 module tend to be outdated or geared toward other displays.
This project serves as a starting point for new users of the Funduino LCM 1602.

Requirements for this projects: 
1. Aduino uno with cable for laptop communication.
2. Funduino LCM 1602 module
3. 4 jumper cables
4. Installation of the Arduino IDE
5. The LiquidCrystalI2C library must be downloaded.

Steps:
1.  Wire your Arduino Uno to your 1602 module as follows:
```
    A4 -> SCL
    A5 -> SDA
    5V -> VCC
    GND -> GND
```

2. Connect to your pc.
3.  Open the Arduino IDE.
4.  If you haven't already, we need to download the LiquidCrystalI2C library.  To do this, go to SKETCH-->INCLUDE A LIBRARY-->MANAGE LIBRARIES
5.  Open up the sketch provided with this repository.
6.  Connect Arduino to the PC.
7.  Verify the correct communication port.  TOOLS --> PORT in the ide
NOTE:  The com port name may very as I understand.  My IDE on my Windows 10 computer recognized COM 3 as the port.
8.  Verify and compile the code SKETCH-->VERIFY/COMPILE.
9.  Save the sketch
10.  Upload the sketch to your arduion SKETCH--> UPLOAD

NOTE:  I have noticed there is a variation between the Arduino IDE on the pi and in Windows.  The versions of each are up to date as of 12/19/2016.
-On the pi, one must dowload new libraries outside of the IDE and then search and import.
-COM ports certainly vary between the pi and windows.
-The pi ide window has a great font-size initially after download, the windos ide needed to be tweaked by going to preferences.
