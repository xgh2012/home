<?php
/**
 * Created by PhpStorm.
 * User: xgh
 * Date: 2019/7/24
 * Time: 11:26
 */

 //人像分割

 //SDK 见百度文档 https://ai.baidu.com/docs#/Body-API/top
require_once 'AipBodyAnalysis.php';

$ocr = new AipBodyAnalysis('16864563','cbtFfv2n0TmPzIV4gT5h3WSD','TWKq2MRGx6TcFOQFXaxCkDHzK7Cxkc5N');

$data = file_get_contents('headimg.jpg');
$res = $ocr->bodySeg($data);
$content = "<?php".PHP_EOL;
$content.= var_export($res,1).";".PHP_EOL;
file_put_contents('xgh.php',$content);

file_put_contents("./baidu/labelmap.png",base64_decode($res['labelmap']));
file_put_contents("./baidu/foreground.png",base64_decode($res['foreground']));
file_put_contents("./baidu/scoremap.png",base64_decode($res['scoremap']));
