package components

//var menuhtml string = `<ul id="menu"><template id="menuitem"><li ><a href="" id="menudata"></a></li></template></ul>`
var menuhtml string = `<template id="menuitem"><li ><a href="" id="menudata"></a></li></template><ul id="menu">this is me</ul>`

var menucss string = `
.tab {
  text-align: center;
}

#menu {
  list-style-type: none;
  padding: 0;
  margin: 0;
  border: 0;
}


  li {
  padding: 10px;
  text-decoration: none;
  display: inline-block;
  background-color: #f4f4f4;
  border-bottom: solid 0px;
  border-top: solid 0px;

  
}

#menu > li:first-child {
  border-left: solid 0px;
  border-bottom-left-radius: 6px;
  border-top-left-radius: 6px;
}

li:last-child {
  border-right: solid 0px;
  border-bottom-right-radius: 6px;
  border-top-right-radius: 6px;
}


a {
  text-decoration: none;
  color:black;
}

li.selected {
  background-color: #f4f4f4;
  font-weight: bold;
 /* border-left: rgb(204, 201, 201) solid;*/
}


  li:hover {

    background-color:grey;

}`
