


function StartRealtime(roomid, timestamp) {
    StartEpoch(timestamp);
    StartSSE(roomid);
    StartForm();
}

function StartForm() {
    $('#chat-message').focus();
    $('#chat-form').ajaxForm(function() {
        $('#chat-message').val('');
        $('#chat-message').focus();
    });
}

function StartEpoch(timestamp) {
    var windowSize = 60;
    var height = 200;
    var defaultData = histogram(windowSize, timestamp);

    window.heapChart = $('#heapChart').epoch({
        type: 'time.area',
        axes: ['bottom', 'left'],
        height: height,
        historySize: 10,
        data: [
            {values: defaultData},
            {values: defaultData}
        ]
    });

    window.mallocsChart = $('#mallocsChart').epoch({
        type: 'time.area',
        axes: ['bottom', 'left'],
        height: height,
        historySize: 10,
        data: [
            {values: defaultData},
            {values: defaultData}
        ]
    });

    window.messagesChart = $('#messagesChart').epoch({
        type: 'time.line',
        axes: ['bottom', 'left'],
        height: 240,
        historySize: 10,
        data: [
            {values: defaultData},
            {values: defaultData},
            {values: defaultData}
        ]
    });
}

function StartSSE(roomid) {
    if (!window.EventSource) {
        alert("EventSource is not enabled in this browser");
        return;
    }
    var source = new EventSource('/stream/'+roomid);