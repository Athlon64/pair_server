<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="static/css/bulma.min.css">
    <title>汉英句库</title>
</head>

<body>
    <section class="hero is-primary is-fullheight">
      <!-- Hero head: will stick at the top -->
      <div class="hero-head">
        <header class="navbar">
          <div class="container">
            <div class="navbar-brand">
              <a class="navbar-item">
                <img src="/static/img/logo.png" width="120" alt="Logo">
              </a>
            </div>
            <div id="navbarMenuHeroC" class="navbar-menu">
              <div class="navbar-end">
                <a class="navbar-item {{if .isHome}} is-active {{end}}" href="/">首页</a>
                <a class="navbar-item {{if .isAll}} is-active {{end}}" href="/all">全部</a>
                <a class="navbar-item {{if .isAdd}} is-active {{end}}" href="/add">添加</a>
              </div>
            </div>
          </div>
        </header>
      </div>
