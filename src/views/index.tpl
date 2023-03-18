{{template "header.tpl" .}}

<!-- Hero content: will be in the middle -->
<div class="hero-body">
    <div class="container has-text-centered">
        <form class="colums" action="/search" method="get">
            <div class="column is-half is-offset-one-quarter">
                <input class="input is-primary is-rounded is-focused" type="text" name="word" placeholder="搜索..." />
            </div>
        </form>
   </div>
</div>

{{template "footer.tpl" .}}
