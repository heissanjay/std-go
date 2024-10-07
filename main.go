package main 

import (
  "fmt"
  "unicode"
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

// J 

// Joins takes a slice of string and return a concatenated string with specified delimiter

//  fruits := []string{"apple", "banana", "orange", "pineapple"}
//  fmt.Println(strings.Join(fruits, "+")) //apple+banana+orange+pineapple 


// L 
// return the last index of the substr in s  
  //fmt.Println(strings.LastIndex("go lang learning", "l"))
// returns the index of any unicode point from chars in s, or -1 if no 
 // fmt.Println(strings.LastIndexAny("go lang Learning", "l")) // output: 3
  /*
    Mostly use Any suffixed method whenever dealing with unicode
  */

 // fmt.Println(strings.LastIndexByte("I know who", 'i')) // -1
 // fmt.Println(strings.LastIndexByte("hello world", 'd')) // 10


 // result := strings.LastIndexFunc("quickbrown@fox", func(r rune)bool{
   // return !unicode.IsNumber(r) && !unicode.IsLetter(r)
 // })

 // fmt.Println(result)

// M 
/*
Map returs a copy of the string s with all its characters modified according
to the mapping functions. If mapping return a negative value, the character
is dropped from the string with no replacement
*/


// Example:

//upperCased := strings.Map(func(r rune)rune{
  //return unicode.ToUpper(r)
//}, "Hello there")

//fmt.Println(upperCased)

// R 

// takes a string, a string to append and it count 

// fmt.Println("ba" + strings.Repeat("na", 2)) //prints: banana


// find and replace kind of

//fmt.Println(strings.Replace("foo bar zoo", "o", "e", 2)) // fee bar zoo 


// Replace all 
//fmt.Println(strings.ReplaceAll("foo bar zoo", "o", "e"))

// S 
//result := strings.Split("Hey my name is John Doe", " ")
//fmt.Println(result)// [Hey my name is John Doe]

//result := strings.SplitAfter("Hey my name is John Doe", " ")
// whenever the instance of the delimiter is encountered, that part 
// will be takes as an item "Hey ", "my ", "name ", "is ", "John ", "Doe"
//fmt.Println(result) // [Hey  my  name  is  John  Doe]
//fmt.Println(len(result))

// split the string into n instance after the delimiter
//result := strings.SplitAfterN("Hey there how are you", " ", 3)
//for i, v := range result {
  //fmt.Printf("Word %v is %v \n", i, v) // len(result) = 3
//}


//result := strings.SplitN("Hey there my name is john cena", " ", 2)

//for i, v := range result {
  //fmt.Printf("item %v is %v\n", i, v)
//}

// T 

//fmt.Println(strings.ToLower("HELLO THERE"))
//fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

// converts each char to upper case 
//fmt.Println(strings.ToTitle("Hello there"))

// result look same as title: TODO should look into how its working differs from Title method 
//fmt.Println(strings.ToUpper("hello world"))


//fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
// takes a string s which may contain invalid utf8 byte sequence
// we can replace those with the replacement string 
// fmt.Println(strings.ToValidUTF8("Hello, \xff\xfe world", "Invalid"))

// capitalize first letter every word 
// fmt.Println(strings.Title("Hello world"))
// removes leading and trailing char 'c' in the string s
//fmt.Println(strings.Trim("hello there", "e")) // output: hello ther 

// trims any leading and trailing numbers
//fmt.Println(strings.TrimFunc("Hello There234", func(r rune)bool{
  //return unicode.IsNumber(r)
//}))


// TrimLeft, Right + Func 
// Trim Prefix, TrimSuffix 
// Trim Space 

//fmt.Println(strings.TrimLeft("Hello world", "Hello worl"))
// remove any leading number 
//fmt.Println(strings.TrimLeftFunc("3245Hello World", func(r rune)bool{
  //return unicode.IsNumber(r)
//}))

// remove any leading and trailing space 
//fmt.Println(strings.TrimSpace("  Hello world  "))
// remove the training substr from s 
// fmt.Println(strings.TrimRight("Hello world", "world"))
//  remove trailing numbers
//fmt.Println(strings.TrimRightFunc("Hello World123", func(r rune)bool{
  //return unicode.IsNumber(r)
//}))

// remove any prefix with the prefix in String s 
//fmt.Println(strings.TrimPrefix("Hello world", "H"))
// removes any suffic with the suffix in string s 
//fmt.Println(strings.TrimSuffix("Hello world", "rld"))

//fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "Önnek İş"))

  /*
  ALL STRING METHOD COVERED
*/ 

}
