// Example starter JavaScript for disabling form submissions if there are invalid fields
(function () {
    'use strict'
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    let forms = document.querySelectorAll('.needs-validation')
    // Loop over them and prevent submission
    Array.prototype.slice.call(forms)
        .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }
                form.classList.add('was-validated')
            }, false)
        })
})();

function notify(msg, msgType) {
    notie.alert({
        type: msgType,
        text: msg,
    })
}

function notifyModal(title, text, icon, confirmationButtonText) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmationButtonText
    })
}

function Prompt() {
    let toast = function (c) {
        const{
            msg = '',
            icon = 'success',
            position = 'top-end',
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
            confirmButtonColor: "#3085d6",
        })

    }

    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
            confirmButtonColor: "#3085d6",
        })

    }

    async function custom(c) {
        console.log(c);
        const {
            icon = "",
            msg = "",
            title = "",
            confirmButtonText = "確定",
            cancelButtonText = "キャンセル"
        } = c
        const { value: result } = await Swal.fire({
            icon: icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            confirmButtonColor: "#3085d6",
            confirmButtonText: confirmButtonText,
            cancelButtonText: cancelButtonText,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined){
                    c.didOpen();
                }
            },
            preConfirm: () => {
                if (c.preConfirm !== undefined){
                    c.preConfirm();
                }
            }
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined) {
                        c.callback(result);
                    }
                } else {
                    c.callback(false);
                }
            } else {
                c.callback(false);
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}

function roomBookConfirm(roomId) {
    let attention = new Prompt();
    let html = `
            <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
                <div id="reservation-dates-modal" class="d-flex content-justify-center p-1">
                   <div class="col-5">
                        <input type="text" name="start" id="start" class="form-control" placeholder="開始日" required autocomplete="off">
                    </div>
                    <div class="col-2">〜</div>
                    <div class="col-5">
                        <input type="text" name="end" id="end" class="form-control" placeholder="終了日" required autocomplete="off">
                    </div>
                </div>
            </form>
            `;
    attention.custom({
        title: '日付を選択してください',
        msg: html,
        didOpen: () => {
            const elem = document.getElementById("reservation-dates-modal");
            const rangepicker = new DateRangePicker(elem, {
                format: "yyyy-mm-dd",
                language: "ja",
                minDate: new Date(),
                orientation: "auto top",
                showOnFocus: false,
            });
        },
        callback: function(result) {
            if (result) {
                let form = document.getElementById("check-availability-form");
                let formData = new FormData(form);
                formData.append("csrf_token", document.getElementById("csrf_token").value);
                formData.append("room_id", roomId)
                console.log(roomId);
                fetch('/search-availability-json', {
                    method: "post",
                    body: formData,
                })
                    .then(response => response.json())
                    .then(data => {
                        console.log(data)
                        if (data.ok) {
                            attention.custom({
                                icon: "success",
                                title: "空きがあります",
                                confirmButtonText: "今すぐ予約",
                                callback: function (result) {
                                    if (result) {
                                        window.location.href = `/book-room?id=${data.room_id}&s=${data.start_date}&e=${data.end_date}`
                                    }
                                }
                            });
                        } else {
                            attention.error({
                                title: "空きがありません",
                                msg: "期間を調整して再度検索してください"
                            });
                        }
                    });
            }
        }
    });
}