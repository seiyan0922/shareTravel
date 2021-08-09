{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="new-event-page">
            <div class="new-event-title">Search Event</div>
            <div class="gray-line"></div>
            <div>
                <div class="new-event-form-box">
                    <form action="/event/search" method="POST">
                        <table class="form-table">
                            <tr>
                                <td>
                                    <label for="auth_key">イベントID</label>
                                </td>
                                <td>
                                    ：<input type="text" name="auth_key" class="form-common">
                                </td>
                            </tr>
                            </table>
                            <div class="common-submit-box">
                                <input class="submit-common" type="submit" value="検索">
                            </div>
                            
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}