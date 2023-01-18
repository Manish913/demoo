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
      $scientificname = $_POST['scientificname'];
      $localname = $_POST['localname'];
      $commonname = $_POST['commonname'];
      $description = $_POST['description'];


     
            //upload an images
            $target_file = "";
            if (isset($_FILES["image"]["name"])) {
                $target_file = "uploads/".basename($_FILES["image"]["name"]);
                $uploadOk = 1;
                $imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));
                // Check if image file is a actual image or fake image
                $check = getimagesize($_FILES["image"]["tmp_name"]);            
                if($check !== false) {
                    move_uploaded_file($_FILES["image"]["tmp_name"], "uploads/" . $_FILES["image"]["name"]);
                    $uploadOk = 1;
                } else {
                    echo "File is not an image.";
                    $uploadOk = 0;
                }
            }
            //end of image upload


        //upload an images
      $target_file = "";
      if (isset($_FILES["image2"]["name"])) {
        $target_file = "../uploads/".basename($_FILES["image2"]["name"]);
        $uploadOk = 1;
        $imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));
        // Check if image file is a actual image or fake image
          $check = getimagesize($_FILES["image2"]["tmp_name"]);      
          if($check !== false) {
            move_uploaded_file($_FILES["image2"]["tmp_name"], "../uploads/" . $_FILES["image2"]["name"]);
              $uploadOk = 1;
          } else {
              echo "File is not an image.";
              $uploadOk = 0;
          }
      }
      //end of image upload

        //upload an images
      $target_file = "";
      if (isset($_FILES["image3"]["name"])) {
        $target_file = "../uploads/".basename($_FILES["image3"]["name"]);
        $uploadOk = 1;
        $imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));
        // Check if image file is a actual image or fake image
          $check = getimagesize($_FILES["image3"]["tmp_name"]);      
          if($check !== false) {
            move_uploaded_file($_FILES["image3"]["tmp_name"], "../uploads/" . $_FILES["image3"]["name"]);
              $uploadOk = 1;
          } else {
              echo "File is not an image.";
              $uploadOk = 0;
          }
      }
      //end of image upload


      try {
          $stmt = $connect->prepare('INSERT INTO addnew (scientificname, localname, commonname, description, image1, image2, image3) VALUES (:scientificname, :localname, :commonname, :description, :image1, :image2, :image3)');
          $stmt->execute(array(
            ':scientificname' => $scientificname,
            ':localname' => $localname,
            ':commonname' => $commonname,
            ':description' => $description,
            ':image1' => $target_file,
            ':image2' => $target_file,
            ':image3' => $target_file,
            ));       

        header('Location: adminaddnewplant.php?action=reg');
        exit;
      }
      catch(PDOException $e) {
        echo $e->getMessage();
      }
  }
 if(isset($_GET['action']) && $_GET['action'] == 'reg') {
    $errMsg = 'Registration successfull. Thank you';
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


    <title>ApnaBagicha</title>
</head>

<body>
    <div class="header sticky-top">
        <nav class="navbar navbar-expand-lg navbar-light " style="background-color:  #0cb637;">
            <div class="container-fluid">
                <a class="navbar-brand" href="#"><b>KNOW YOUR GARDEN</b></a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                    aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="navbar-collapse collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                            <a class="nav-link " aria-current="page" href="dashboard.php">Dashboard</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="#">Add New Plants</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#">View Feedback</a>
                        </li>
                    </ul>

                    <div class="mx-2">
                        <div class=" my-4 button-container col-6">
                            <a href="index.php" class="btn btn-outline-primary">Logout</a>
                        </div>

                    </div>
                </div>
            </div>
        </nav>
    </div>



    <div class="container my-4">
        <div class=" my-4 col-md-7">
            <h1><b>Add the plants</b></h1>
        </div>

        <form class="row g-3" action="" method="post">
            <div class="col-md-6">
              <label for="inputplant" class="form-label">Plant Id</label>
              <input type="int" class="form-control" id="inputplant" name="plantid">
            </div>
            <div class="col-md-6">
              <label for="inputcname" class="form-label">Plant Common Name</label>
              <input type="text" class="form-control" id="inputcname" name="commonname">
            </div>
            <div class="col-md-6">
                <label for="inputsname" class="form-label">Scientific Name</label>
                <input type="text" class="form-control" id="inputsname" name="scientificname">
              </div>
              <div class="col-md-6">
                <label for="inputlname" class="form-label">Local name</label>
                <input type="text" class="form-control" id="inputlname" name="localname">
              </div>
            <div class="col-12">
                <label for="Textarea1" class="form-label">Plant Details</label>
                <textarea class="form-control" id="Textarea1" rows="10" name="description"></textarea>
            </div>
            <div class="col-md-4">
                <label for="formFile" class="form-label">Upload Image 1</label>
                <input class="form-control" type="file" id="formFile" name="image">
              </div>
              <div class="col-md-4">
                <label for="formFile2" class="form-label">Upload Image 2</label>
                <input class="form-control" type="file" id="formFile2" name="image2">
              </div>
             <div class="col-md-4">
                <label for="formFile2" class="form-label">Upload Image 3</label>
                <input class="form-control" type="file" id="formFile2" name="image3">
              </div>
            <div class="col-12">
              <button type="submit" class="w-100 btn btn-primary" name="register" value="register">Add Plant</button>
            </div>
          </form>
    </div>

    <hr class="featurette-divider">
    <footer class="container">
        <p class="float-end"><a href="#">Back to top</a></p>
        <p>© 2022 Know your Garden, pvt. · <a href="#">Privacy</a> · <a href="#">Terms</a></p>
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