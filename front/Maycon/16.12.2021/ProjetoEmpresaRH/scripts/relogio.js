function init(){
    clock();
    setInterval(clock, 1000);
}

function clock(){
    var now = new Date();
    var ctx = document.getElementById('relogio').getContext('2d');

    ctx.save();
    ctx.clearRect(0, 0, 150, 150);
    ctx.translate(75, 75);
    ctx.scale(0.4, 0.4);
    ctx.rotate(-Math.PI/2);
    ctx.strokeStyle = "black";
    ctx.fillStyle = "white";
    ctx.lineWidth = 8;
    ctx.lineCap = "round";

    //hour marks
    ctx.save();
    for (var i =0; i <12; i++){
        ctx.beginPath();
        ctx.rotate(Math.PI / 6);
        ctx.moveTo(100,0);
        ctx.lineTo(120,0);
        ctx.stroke();
    }
    ctx.restore();

    //minute marks
    ctx.save();
    ctx.lineWidth = 5;
    for (i = 0; i < 60; i++){
        if (i % 5 != 0){
            ctx.beginPath();
            ctx.moveTo(117,0);
            ctx.lineTo(120,0);
            ctx.stroke();
        }
        ctx.rotate(Math.PI / 30)
    }
    ctx.restore();
    var sec = now.getSeconds();
    var min = now.getMinutes();
    var hr = now.getHours();
    hr = hr => 12 ? hr - 12 : hr;

    ctx.fillStyle="black";

    // write Hours

    ctx.save();
    ctx.rotate(hr*(Math.PI / 6) + (Math.PI/360) * min + (Math.PI / 21600) * sec);
    ctx.lineWidth = 14;
    ctx.beginPath();
    ctx.moveTo(-20,0);
    ctx.lineTo(80,0);
    ctx.stroke();
    ctx.restore();
}