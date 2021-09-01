{{ define "sidebar" }}
<div class="side-bar">
    <a href="/event/show?event_id={{.Event.Id}}" class="member-add">
        <div class="icon-common">
            <div class="jet-icon"></div>
            <div class="add-member-text">イベントTOP</div>
        </div>
    </a>
    <a href="/member/add?event_id={{.Event.Id}}" class="member-add">
        <div class="icon-common">
            <div class="member-icon"></div>
            <div class="add-member-text">メンバー追加</div>
        </div>
    </a>
    <a href="/expense/add?event_id={{.Event.Id}}" class="member-add">
        <div class="icon-common">
            <div class="expense-icon"></div>
            <div class="expense-text">会計追加</div>
        </div>
    </a>
    <div class="icon-common">
        <a href="/event/download?event_id={{.Event.Id}}" class="member-add">
            <div class="download-icon"></div>
            <div class="download-text">ダウンロード</div>
        </a>
    </div>
</div>
{{end}}