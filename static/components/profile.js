new Vue({
    el: "#app",
    data: {
        weight:71,
        height:170,
        gender:"male",
        user:{id:0, points:0, username:"", level:""},
    },
    created(){
        setTimeout(()=> {
            this.user=User;
        }, 1000)
    }

});