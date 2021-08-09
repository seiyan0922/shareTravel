{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="new-event-page">
            <div class="new-event-title">New Event</div>
            <div class="gray-line"></div>
            <div>
                <div class="new-event-form-box">
                    <form action="/event/confirm" method="POST">
                        <table class="form-table">
                            <tr>
                                <td><label for="name">イベント名</label></td>
                                <td>：<input type="text" name="name" class="form-common"></td>
                            </tr>
                            <tr>
                                <td><label for=date>開催日</label></td>
                                <td>：<input type="date" name="datetime" class="form-common"><td>
                            </tr>
                            

                        </table>
                        <div class="common-submit-box">
                            <input class="submit-common" type="submit" value="確認">
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}