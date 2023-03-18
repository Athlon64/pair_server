{{template "header.tpl" .}}

<div class="hero-body">
    <div class="container has-text-centered">
            <p class="title">找到 <span class="has-text-info">{{.number}}</span> 条结果：</p>
            {{range $index, $pair := .results}}
                 <div class="box"> 
                      <div class="notification is-primary"> 
                           <div class="content is-large">{{$pair.chs}}</div><br> 
                           <div class="content is-large">{{$pair.en}}</div><br> 
                           <a class="button is-danger" href="/del/{{$pair.id}}">删除</a> 
                      </div> 
                 </div> 
            {{end}}
    </div>
</div>

{{template "footer.tpl" .}}
