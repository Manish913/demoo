<?php
   require 'C:\xampps\htdocs\ApnaBagicha\config.php';
   if(isset($_POST['login'])) {

    // Get data from FORM
    $username = $_POST['username'];
    $password = $_POST['password'];

    try {
      $stmt = $connect->prepare('SELECT * FROM user WHERE username = :username');
      $stmt->execute(array(
        ':username' => $username,
        ));
      $data = $stmt->fetch(PDO::FETCH_ASSOC);

      if($data == false){
        $errMsg = "User $username not found.";
      }
      else {
        if(md5($password) == $data['password']) {

          header('Location: dashboard.php');
            echo "Login Successfully";
          exit;
        }
        else
          $errMsg = 'Password not match.';
      }
    }
    catch(PDOException $e) {
      $errMsg = $e->getMessage();
    }
  }
?>
<!doctype html>
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
                    <a href="register.php" class="btn btn-outline-primary">Register</a>
                </div>
                
            </div>
        </div>
    </nav>
</div>
    




    <div class="text-center container my-4">
        <main class="form-signin">
            <form action="" method="post">
              <img class="mb-4" src="login.jpeg" alt="" width="72" height="57">
              <h1 class="h3 mb-3 fw-normal">Please sign in</h1>
          
              <div class="my-2 form-floating">
                <input type="text" class="form-control" id="floatingInput" placeholder="name@example.com" name="username">
                <label for="floatingInput">Email address</label>
              </div>
              <div class="form-floating my-2">
                <input type="password" class="form-control" id="floatingPassword" placeholder="Password" name="password">
                <label for="floatingPassword">Password</label>
              </div>
          
              <div class="checkbox mb-3">
                <label>
                  <input type="checkbox" value="remember-me"> Remember me
                </label>
              </div>
              <button class="w-100 btn btn-lg btn-primary" type="submit" name="login" value="login">Sign in</button>
             
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