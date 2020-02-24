# GolangExercise
The exercise for Golang


## Amtran Design - Sony SN Covert ##

**Sony SN Covert** is a machine for producing serial number.
###[Serial Number Rule]###
---
![Rule Image](/sony_sn_rule1.jpg)


###[User Manual]###
---
##### Download Code #####

1.git clone ssh://fwroot@192.168.1.84/home/fwroot/SB\_proj\_git/sony\_sn\_convertor

#####Run .exe#####

1.run win\_program\_gen.bat  
2.run dist/main.exe

###[Programmer Manual]###
---
#####Enviroment #####
1.[Python3.8.1](https://www.python.org/downloads/)  
2.[PyQt5](https://www.riverbankcomputing.com/software/pyqt/download5)  
3.[Qt Designer](https://build-system.fman.io/qt-designer-download )  
4.To install pyinstaller by the following methods:  
`pip install pyinstaller`  

#####Code #####


1.main.py - 
<font color=red>
Main control UI functions
</font>  
2.amtran\_magic\_box.py - 
<font color=red>
Main class to process all behavior about serial number and file produced
</font>  
3.amtran\_ui\_design.py -
<font color=red>
UI define class
</font>  
4.sony.ui - 
<font color=red>
UI project for Qt Designer
</font>  
5.images.qrc -
<font color=red>
Image resource
</font>  
6.images\_qr.py -
<font color=red>
Image resource table from <font color=blue>5.</font>
</font> 

#####Build .ui to .py(Class) #####
1.`pyuic5 sony.ui -o amtran_ui_design.py`

>**Do Not Modify <font color=red>"amtran\_ui\_design.py"</font>  
>All your action need to be designed in<font color=red> "main.py"</font>**
#####Compile and Execute#####
1.`python main.py`  - 
<font color=red>
Complie .py code
</font>   
2.`pyinstaller -F -w ./main.py` - 
<font color=red>
Turn "main.py" code into "main.exe"
</font>   
