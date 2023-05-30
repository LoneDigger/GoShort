const reg = new RegExp(/^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$/);

$(document).ready(function () {

  toastr.options = {
    "closeButton": false,
    "debug": false,
    "newestOnTop": false,
    "positionClass": "toast-bottom-right",
    "preventDuplicates": false,
    "onclick": null,
    "showDuration": "500",
    "hideDuration": "1000",
    "timeOut": "5000",
    "extendedTimeOut": "3000",
    "showEasing": "swing",
    "hideEasing": "linear",
    "showMethod": "fadeIn",
    "hideMethod": "fadeOut"
  }

  $("#submit").click(function () {
    $("#short_url").val("");

    const url = $("#long_url").val();

    if (!reg.test(url)) {
      toastr.warning("Please enter url.");
      return;
    }

    const json = JSON.stringify({
      "url": url,
    });

    axios.post('/', json, {
      timeout: 5000,
    },
    ).then(function (response) {
      switch (response.data.code) {
        case 0:
          $("#short_url").val(window.location.origin + "/" + response.data.url);
          break;

        case 1:
          toastr.error("Oops!");
          break;

        case 2:
          toastr.warning("Please enter url.");
          break;
      }
    }).catch(function (error) {
      toastr.error(error.code);
    });
  });
});
