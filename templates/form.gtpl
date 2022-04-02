<html>
    <head>
        <title>Form</title>
    </head>
    <body>
        <h3>Form Sample</h3>
        <form action="/submit" method="get">
            <p>param1 <input type="text" name="param1"></p>
            <p>param2 <input type="password" name="param2"></p>
            <input type="submit" value="submit(get)">
        </form>
        <form action="/submit" method="post">
            <p>param1 <input type="text" name="param1"></p>
            <p>param2 <input type="password" name="param2"></p>
            <input type="submit" value="submit(post)">
        </form>
    </body>
</html>