{{template "header.tpl" .}}

<div class="hero-body">
    <div class="container has-text-centered">
        <form action="/doAdd" method="post">
            <input class="input is-primary is-rounded" type="text" name="chs" placeholder="中文"/> <br><br>
            <input class="input is-primary is-rounded" type="text" name="eng" placeholder="英文" /> <br><br>
            <input class="button is-primary is-link" type="submit" value="添加">
        </form>
	</div>
</div>

{{template "footer.tpl" .}}
