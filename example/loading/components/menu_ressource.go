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
  border-bottom: solid 0px;
  border-top: solid 0px;

  
}



a {
  text-decoration: none;
  color:black;
}

li.selected {
  font-weight: bold;
}`
