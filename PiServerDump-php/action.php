<!DOCTYPE html>
<html>
<head>
<title>ReturnUserInfo</title>
<style type="text/css" >
  p {
    font-weight: bold;}
</style>
</head>
<body>
  
<p><a href="index.html">HOME</a></p>
<p>Hi <?php echo htmlspecialchars($_POST['name']); ?>.
You are <?php echo (int)$_POST['age']; ?> years old.</p>
</body>
</html>
