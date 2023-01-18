<?php
  require 'C:\xampps\htdocs\ApnaBagicha\config.php';

  use PHPMailer\PHPMailer\PHPMailer;
  use PHPMailer\PHPMailer\Exception;

  require 'PHPMailer/src/Exception.php';
  require 'PHPMailer/src/PHPMailer.php';
  require 'PHPMailer/src/SMTP.php';

  if(isset($_POST['register'])) {
    $errMsg = '';

    // Get data from FROM
    $fullname = $_POST['fullname'];
    $username = $_POST['username'];
    $email = $_POST['email'];
    $password = $_POST['password'];
    

  
      try {
        $stmt = $connect->prepare('INSERT INTO user (fullname,  email, username, password) VALUES (:fullname, :email, :username, :password)');
        $stmt->execute(array(
          ':fullname' => $fullname,
           ':email' => $email,
          ':username' => $username,
          ':password' => md5($password),
         
          
          ));
        header('Location: Register.php?action=joined');
        exit;
      }
      catch(PDOException $e) {
        echo $e->getMessage();
      }
  }

  if(isset($_GET['action']) && $_GET['action'] == 'joined') {
    $errMsg = 'Registration successfull. Now you can login';
  }
?>


<html lang="en">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet"
        integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
    <link rel="stylesheet" href="index.css">

    <title>ApnaBagicha</title>
</head>

<body>
    <div class="header sticky-top">
    <nav class="navbar navbar-light " style="background-color:  #0cb637;">
        <div class="container-fluid">
            <a class="navbar-brand"><b>KNOW YOUR GARDEN</b></a>
            <div class="mx-2">
                <div class=" my-4 button-container col-6">
                    <a href="login.php" class="btn btn-outline-primary">Login</a>
                </div>
                
            </div>
        </div>
    </nav>
</div>
    




    <div class="text-center container my-4">
        <main class="form-signin">
            <form action="" method="post">
              <img class="mb-4" src="signuplogo.png" alt="" width="72" height="57">
              <h1 class="h3 mb-3 fw-normal">Please Register Here</h1>

              <div class="my-2 form-floating">
                <input type="text" class="form-control" id="floatingInput" placeholder="name" name="fullname" required>
                <label for="floatingInput">Full Name</label>
              </div>

             
              <div class="my-2 form-floating">
                <input type="email" class="form-control" id="floatingInput" placeholder="name@example.com" name="email" required>
                <label for="floatingInput">Email address</label>
              </div>
               <div class="my-2 form-floating">
                <input type="text" class="form-control" id="floatingInput" placeholder="userid" name="username" required>
                <label for="floatingInput">User-Id</label>
              </div>

              <div class="form-floating my-2">
                <input type="password" class="form-control" id="floatingPassword" placeholder="Password" name="password" required>
                <label for="floatingPassword">Password</label>
              </div>
            
              <button class="w-100 btn btn-lg btn-primary" type="submit" name="register" value="register">Register</button>
             
            </form>
          </main>
    </div>

    <hr class="featurette-divider">
    <footer class="container">
        <p class="float-end"><a href="#">Back to top</a></p>
        <p>© 2022 Know your Garden, pvt <a href="#">Privacy</a> · <a href="#">Terms</a></p>
    </footer>







    <!-- Optional JavaScript; choose one of the two! -->

    <!-- Option 1: Bootstrap Bundle with Popper -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM"
        crossorigin="anonymous"></script>

    <!-- Option 2: Separate Popper and Bootstrap JS -->
    <!--
    <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js" integrity="sha384-IQsoLXl5PILFhosVNubq5LC7Qb9DXgDA9i+tQ8Zj3iwWAwPtgFTxbJ8NT4GN1R8p" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.min.js" integrity="sha384-cVKIPhGWiC2Al4u+LWgxfKTRIcfu0JTxR+EQDz/bgldoEyl4H0zUF0QKbrJ0EcQF" crossorigin="anonymous"></script>
    -->
</body>

</html>