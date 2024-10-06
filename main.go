package main 

import (
  "fmt"
  "strings"
)

/*
  learning go strings package
  C (10)
    - Clone
    - Compare
    - Contains
    - ContainsAny
    - ContainsFunc
    - ContainsRune
    - Count
    - Cut
    - CutPrefix
    - CutSuffix
  
  E (1)
    - EqualFold

  F (2)
    - Fields
    - FieldsFunc
    
  H(2)
    - HasPrefix
    - HasSuffix

  I (5)
   - Index
   - IndexAny
   - IndexByte
   - IndexFunc
   - IndexRune
   
  J (1)
   -Join

  L(4)
   - LastIndex
   - LastIndexAny
   - LastIndexBytes
   - LastIndexFunc
  
  M(1)
   - Map

  R(3)
   - Repeat
   - Replace
   - ReplaceAll
  
  S(5)
   - Split
   - SplitAfter
   - SplitAfterN
   - SplitN

  T
   - Title
   - ToLower
   - ToLowerSpecial
   - ToTitle
   - ToTitleSpecial
   - ToUpper
   - ToUpperSpecial
   - ToValidUTF8
   - Trim
   - TrimFunc
   - TrimLeft
   - TrimLeftFunc
   - TrimPrefix
   - TrimRight
   - TrimRightFunc
   - TrimSpace
   - TrimSuffic

   All others are types
   - Builder
   - Reader
   - Replacer


Are we going to write example / try out each of this function?
"YES", let's go wild;
*/

func main(){

  // "C"
  
//  original := "I'm the OG"
  //cloned := strings.Clone(original)
  //fmt.Println(cloned) // No you are not
  
  // lets compare the og with clone
  //fmt.Println(strings.Compare(original,cloned)) // returns 0.

  // contains, containsAny, containsFunc
  //parent := "I have house, car, bike"
  //children := "I have house"

//  fmt.Println(strings.Contains(parent, children))
  //checks if the word contains any characters
  //vowels := "aeiou"
  //word := "running"

  //fmt.Println(strings.ContainsAny(word, vowels)) //output: true

  // ContainsFunc takes a additional func which validates any condition on the string

  // let validate a email

//  email := "example@example.com"
//  fmt.Println(strings.ContainsFunc(email, func(r rune)bool{
//    return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '_' || r == '@'
//  }))

  // Count actually counts the occurance of the character/string

//  fmt.Println(strings.Count("Hello World", "l")) // output: 3
//  fmt.Println(strings.Count("Hello World", "World")) // output: 1

//    fmt.Println(strings.ContainsRune("Hello world", 'r'))

    
//    before, after, found := strings.Cut("Hello world", "l") // split the word with the first occ of the substr
    
//fmt.Println(before)
  //  fmt.Println(after)
   // fmt.Println(found)

 // result, _ := strings.CutPrefix("Hello world","Hello")
 // fmt.Println(result) // " world"

 // result, _ := strings.CutSuffix("Hello world", "world")
//  fmt.Println(result)// "Hello "

/*
    ALL C are Done
*/
    
  /*
    -  Are are belongs to same encoding -> are the characters same (ignore-case), if both
      YES, then true otherwise false
  */
// E
 // fmt.Println(strings.EqualFold("ß", "ß")) // return true

// F
  
 // splits the string s around each instance of one or more consecutive white space character.
 /*
message := "Hey, My name is John Doe"
result := strings.Fields(message)
fmt.Println(len(result))

for _, v := range result {
  fmt.Println(v)
}
*/

/*

  FieldFunc takes a additional function and store value that statisfies the given function 
*/

 // fmt.Println(strings.FieldsFunc(" foo1;bar2,bar3...", func(r rune)bool{
 // return !unicode.IsLetter(r) && !unicode.IsNumber(r)
 // }))
  
  // H
// check whether the string has the given prefix or not 
 // fmt.Println(strings.HasPrefix("Hello world", "hello")) // return: False
  // check whether the string has the given suffix or not
//  fmt.Println(strings.HasSuffix("Hello world", "world")) // return: True

// I
// Index return the first occ of the give substring
//  fmt.Println(strings.Index("Hello world", "e")) // outputs: 1
  
//  fmt.Println(strings.IndexAny("Hello world", "hello")) // outputs: 1 (e,e)
  
//  fmt.Println(strings.IndexByte("golang", 'g'))
//  fmt.Println(strings.IndexByte("unicode", 'o'))



/*
  IndexFunc return the index of the first rune in string s, that return true for the function func f 
*/

//result := strings.IndexFunc("Hello world", func(r rune)bool{
//return unicode.IsNumber(r) 
//}) // return -1 as no digits found
//fmt.Println(result)

// IndexRune takes a string and rune and the index of the first occ of the rune in string s
// fmt.Println(strings.IndexRune("Hello", 'e'))





}
