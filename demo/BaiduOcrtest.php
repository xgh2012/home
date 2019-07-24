<?php
/**
 * Created by PhpStorm.
 * User: xgh
 * Date: 2019/7/24
 * Time: 11:26
 */
require_once 'AipOcr.php';

$ocr = new AipOcr('','','');

$data = file_get_contents('zhengmian.jpg');
$res = $ocr->idcard($data,'front',['detect_direction'=>true]);
$content = "<?php".PHP_EOL;
$content.= var_export($res,1).";".PHP_EOL;
file_put_contents('AipOcrRes.php',$content);
die;