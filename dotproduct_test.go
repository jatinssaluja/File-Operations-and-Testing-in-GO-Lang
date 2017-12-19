package main

import (
  
  "testing"
  "os"
  "io/ioutil"
  
)


func TestCreateMatrices(t *testing.T){
  
  fileData1:= []string{"1 2 3","4 5 6"}
  fileData2:= []string{"7 8","9 10","11 12"}
  
  matrix1, matrix2:= CreateMatrices(fileData1, fileData2)
  
  if (!(len(matrix1) ==2 && len(matrix1[0]) ==3 && len(matrix2) ==3 && len(matrix2[0]) ==2 )) {
   
      t.Errorf("Matrices Not Created Correctly")
    
  }
  
  
  
}


func TestDotProduct(t *testing.T){
  
  fileData1:= []string{"1 2 3","4 5 6"}
  fileData2:= []string{"7 8","9 10","11 12"}
  
  matrix1, matrix2:= CreateMatrices(fileData1, fileData2)
  
  matrix3:= DotProduct(matrix1,matrix2)
  
  if len(matrix3) != len(matrix1){
    
     t.Errorf("Expected the number of rows in the result matrix to be %v , but got %v", len(matrix1), len(matrix3))
    
    }
  
  if len(matrix3[0]) != len(matrix2[0]){
  
     t.Errorf("Expected the number of columns in the result matrix to be %v , but got %v", len(matrix2[0]),len(matrix3[0]))
    
    }
  
  if matrix3[0][0] !=58 {
    t.Errorf("Expected 58 at matrix3[0][0], but got %v", matrix3[0][0])
    
  }
  
  if matrix3[0][1] !=64 {
    t.Errorf("Expected 64 at matrix3[0][1], but got %v", matrix3[0][1])
    
  }
  
  if matrix3[1][0] !=139 {
    t.Errorf("Expected 139 at matrix3[1][0], but got %v", matrix3[1][0])
    
  }
  
  if matrix3[1][1] !=154 {
    t.Errorf("Expected 154 at matrix3[1][1], but got %v", matrix3[1][1])
    
  }
  
  
}


func TestReadFromFile(t *testing.T){

  os.Remove("_demotesting")
  
  fileData:= "1\n2\n3"
  
  ioutil.WriteFile("_demotesting",[]byte(fileData),0666)
  
  loadedFileData:= ReadFromFile("_demotesting")
  
  if len(loadedFileData)!=3 {
    
     t.Errorf("Expected 3 lines from the file, but got %v", len(loadedFileData))
    
  }
  
  os.Remove("_demotesting")
  
}










