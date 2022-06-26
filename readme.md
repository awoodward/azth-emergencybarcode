# Emergency Barcode Generator
Command line tool for generating emergency files from a template file. 

Run from terminal or command window. 

```
usage: ./emergencybarcode

Program for generating emergencies formatted for pressure seal forms 
for Arizona Treasure Hunt
Written by: Andy Woodward - TH Committee 2019-2022
Email: awoodward@gmail.com

Options:
  -carCount int
    	Number of cars (default 100)
  -file string
    	Destination file for single emergency file (default "emergencies.pdf")
  -noBarcode
    	Disable printing of barcodes
  -noSecurity
    	Disable printing of security hashing
  -pages int
    	Number of pages in template (default 52)
  -singleFile
    	Combine all into one file
  -startingEmergency int
    	First emergency in the sequence (default 1)
  -template string
    	Emergency template file (default "template.pdf")
```



Example:

```
./emergencybarcode -carCount 95 -pages 52 -template “template file.pdf”
```

Generates individual emergency files for each car from a 52-page template file named “template file.pdf.”

```
./emergencybarcode -carCount 95 -pages 2 -template “template file.pdf” -singleFile -startingEmergency 5
```

Generates a single emergency file from a 2-page template file named “template file.pdf” starting at emergency number 5. This is used to reprint an emergency if there are corrections. 
