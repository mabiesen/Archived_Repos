<!DOCTYPE html>
<html>
<!-- This is an HTML file containing things I like to keep up on -->
  <head>
  <meta http-equiv="refresh" content="30">
    <title>MyStuff</title>
    <style type="text/css">
  html {
    background-color: #FAFAD2;}
  div.main_page {
    background-color: rgb(255,255,255);
    position: relative;
    display: table;

    width: 800px;

    margin-bottom: 3px;
    margin-left: auto;
    margin-right: auto;
    padding: 0px 0px 0px 0px;

    border-width: 2px;
    border-color: #212738;
    border-style: solid;}

      body {
          font-family: arial;
	  font-size: 20px;
          background-color: #FAFAD2;}

      h1 {
          color: rgb(0,0,0);
          font-size: 100px;
          margin: 10px 10px 10px 10px;
          }
      a {
	  color: rgb(255,0,0);
          font-weight: bold;}

      div.winner {
	  font-size: 50px;}

      table.second {
          float: left;
          margin: 0px 5px 5px 5px;
          border: 2px solid black;}
      table.first {
          float: left;
	  margin: 0px 5px 5px 5px;
	  border: 2px solid black}
      th {
          text-decoration: underline;}
      td, th {
          padding: 5px 5px 5px 5px;}
      th.one {
          text-decoration: none;
          font-size: 30px;}
      .clear {
	  clear: left;}
    </style>
  </head>
  <div class="main_page">
  <body>
    <h1>MyStuff</h1>
    <h2>MyLinks</h2>
    <p>Below you will find some links to the websites I commonly use:</p>
    <table class="first">
      <tr>
	<th scope="col" colspan="2" class="one">Links for Learning Code</th>
      <tr>	
        <th scope="col">Type</th>
	<th scope="col">Link</th>
      <tr>
	<td>GuidedLearning</td>
	<td><a href="https://www.codecademy.com">CodeCademy</a></td>
      </tr>
      <tr>
        <td>CodePuzzles</td>
        <td><a href="https://www.reddit.com/r/dailyprogrammer">Reddit</a></td>
      </tr>
      <tr>
        <td>CommunityHelp</td>
        <td><a href="https://www.stackoverflow.com">StackOverflow</a></td>
      </tr>
      <tr>
	<td>CodeStorage</td>
        <td><a href="https://www.github.com">GitHub</a></td>
      </tr>
      <tr>
	<td>BasicSyntax
	<td><a href="https://www.cheat-sheets.org">CheatSheets</td>
      </tr>
    </table>

<table class="second">
      <tr>
        <th scope="col" colspan="2" class="one">Links that I Use Daily</th>
      <tr>
        <th scope="col">Type</th>
        <th scope="col">Link</th>
      <tr>
        <td>News</td>
        <td><a href="http://www.drudgereport.com">DrudgeReport</a></td>
      </tr>
      <tr>
        <td>News</td>
        <td><a href="https://www.shtfplan.com">SHTFPlan</a></td>
      </tr>
      <tr>
	<td>Email</td>
	<td><a href="https://www.outlook.com">Outlook</a></td>
      </tr>
      <tr>
	<td>Email</td>
	<td><a href="https://www.gmail.com">GMail</a></td>
      </tr>
      <tr>
	<td>Google</td>
	<td><a href="https://www.google.com">Google</a></td>
      </tr>
</table>
<p class="clear"><hr></p>
<h2>MyTrackers</h2>
<p>Below you will find a few items of interest that I track:</p>
<?php ;
  echo "The information below is valid as of " . date("h:i:sa")." on ".date("m-d-Y");
?> 
<h3>PreciousMetals</h3>
<?php include("getmystuff.php");
  echo "The price of gold is " . $getgold;
?>
<?php ;
  echo "The price of silver is " . $getsilver;
?>
<h3>Weather</h3>
<p>The temperature as of ____ is ____</p>
<p>The humidity as of ____ is ____</p>
<p>The chance of rain as of ____ is ____</p>

<h3>Electoral Votes</h3>

<?php ;
  echo "<p>Hillary has :".$hillary." electoral votes</p>";
  echo "<p>Donald has :".$donald." electoral votes</p>";
?>
</body>
</div>
</html>
