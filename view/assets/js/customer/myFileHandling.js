function gettingFileHandling() {
var max_fields = 6;
var wrapper = document.getElementById("filesArea");
var add_button = document.getElementById("add_form_field");

var x = 1;
$(add_button).click(function(e) {

    e.preventDefault();
    if (x < max_fields) {
        x++;
        $(wrapper).append('<div class="col-sm-6"> <input type="file"name="file'+x+'" class="form-control " id="file" placeholder="Project Name" required> </div> <div class="col-sm-6"> <button type="button" id="add_form_field" class="btn btn-primary  btn-block">Delete </button> <hr></div> '); //add input box
    } else {
        alert('You Reached the limits')
    }
});

$(wrapper).on("click", ".delete", function(e) {
    e.preventDefault();
    $(this).parent('div').remove();
    x--;
});

}