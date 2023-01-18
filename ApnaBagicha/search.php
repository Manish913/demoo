<?php
  require 'C:\xampps\htdocs\ApnaBagicha\config.php';

  $data = [];
  
  if(isset($_POST['search'])) {
    // Get data from FORM
    $keywords = $_POST['keywords'];

    //keywords based search
    $keyword = explode(',', $keywords);
    $concats = "(";
    $numItems = count($keyword);
    $i = 0;
    foreach ($keyword as $key => $value) {
      # code...
      if(++$i === $numItems){
         $concats .= "'".$value."'";
      }else{
        $concats .= "'".$value."',";
      }
    }
    $concats .= ")";
  //end of keywords based search
  
    
    try {
      //foreach ($keyword as $key => $value) {
        # code...

        $stmt = $connect->prepare("SELECT * FROM addnew WHERE plantid IN $concats" );
        $stmt->execute();
        $data2 = $stmt->fetchAll(PDO::FETCH_ASSOC);

       // $stmt = $connect->prepare("SELECT * FROM addnew WHERE plantid IN $concats");
       // $stmt->execute();
        $data8 = $stmt->fetchAll(PDO::FETCH_ASSOC);

        $data = array_merge($data2, $data8);

    }catch(PDOException $e) {
      $errMsg = $e->getMessage();
    }
  }


  if(isset($_POST['comment'])) {
      $errMsg = '';
      // Get data from FROM
      $plantid = $_POST['plantid'];
      $purposes = $_POST['purposes'];
      $other = $_POST['other'];
      $Comments= $_POST['Comments'];


      try {
          $stmt = $connect->prepare('INSERT INTO feedback (plantid, purposes, other, Comments) VALUES (:plantid, :purposes, :other, :Comments)');
          $stmt->execute(array(
            ':plantid' => $plantid,
            ':purposes' => $purposes,
            ':other'=> $other,
            ':Comments' => $Comments,
            ));       

        header('Location: search.php?action=reg');
        exit;
      }
      catch(PDOException $e) {
        echo $e->getMessage();
      }
  }
 if(isset($_GET['action']) && $_GET['action'] == 'reg') {
    $errMsg = 'Your Feedbacks are very important for us. Thank you';
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
                    <a href="login.php" class="btn btn-outline-primary">Login</a>
                </div>
                
 

            </div>
        </div>
    </nav>
</div>
    




    <div class="container my-4">
        <div class="my-4 col-md-7">
            <h1><b>Find Here Your Curiosity</b></h1>
        </div>
        <form class="my-4 d-flex" action="" method="post">
            <input class="form-control me-2" type="search" placeholder="Search Trees or Plants by Id" aria-label="Search" name="keywords">
            <button class="btn btn-outline-success" type="submit" name="search" value="search">Search</button>
          </form>
          <?php
              if(isset($errMsg)){
                echo '<div style="color:#FF0000;text-align:center;font-size:17px;">'.$errMsg.'</div>';
              }
              if(count($data) !== 0){
                echo "<h2 class='text-center'>Your result</h2>";
              }else{
                //echo "<h2 class='text-center' style='color:red;'>Try Some other keywords</h2>";
              }
            ?>        
            <?php 
                foreach ($data as $key => $value) {           
                  echo '<div class="card card-inverse card-info mb-3" style="padding:1%;">          
                        <div class="card-block">';
                          // echo '<a class="btn btn-warning float-right" href="update.php?id='.$value['id'].'&act=';if(isset($value['ap_number_of_plats'])){ echo "ap"; }else{ echo "indi"; } echo '">Edit</a>';
                         echo   '<div class="row">
                            <div class="col-4">
                            <h4 class="text-center">Plant Details</h4>';
                              echo '<p><b>PlantID: </b>'.$value['plantid'].'</p>';
                              echo '<p><b>scientificname: </b>'.$value['scientificname'].'</p>';
                              echo '<p><b>localname: </b>'.$value['localname'].'</p>';
                              echo '<p><b>commonname: </b>'.$value['commonname'].'</p>';
                              echo '<p><b>description: </b>'.$value['description'].'</p>';
                              if ($value['image1'] !== 'C:\xampps\htdocs\Bagicha\uploads/') {
                                # code...
                                echo '<img src="C:\xampps\htdocs\Bagicha\uploads/'.$value['image1'].'" width="100">';
                              }
                              if ($value['image2'] !== 'C:\xampps\htdocs\Bagicha\uploads/') {
                                # code...
                                echo '<img src="C:\xampps\htdocs\Bagicha\uploads/'.$value['image2'].'" width="100">';
                              }
                             
                            echo '</div>
                          </div>  
                          <form action="" method="POST">
                          <table width=100%><tr>
                          <th>
                          <h3> Pupose of visiting</h3></th><th><h3>Your Comments</h3></th></tr><br><tr><td>
                          <input type"text" name="plantid" placeholder="Enter PlantID" ><br>
                          <input type="radio" id="Medicinal" name="purposes" value="Medicinal">
                          <label for="Medicinal">Medicinal</label><br>
                          <input type="radio" id="Knowledge" name="purposes" value="Knowledge">
                          <label for="Knowledge">Knowledge</label><br>
                          <input type="radio" id="viewport" name="purposes" value="Viewing">
                          <label for="Viewing">Viewing</label></br>
                          <input type="radio" id="viewport" name="purposes" value="other">
                          <label for="Other">Other</label></br>
                          <input type="text" name="other" Placeholder="Other Purposes"><br>
                          </td>
                          <td>
                          <input type="text" name="Comments" placeholder="your Comments" style="height:200px;"><br>
                          <button type="submit" name="comment">Submit</button>  </form>   
                          </td>
                          </tr> 
                          </table>      
                         </div>
                      </div>';
                }
              ?>             
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