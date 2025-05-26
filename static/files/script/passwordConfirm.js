function eot(id) {
  const elem = document.getElementById(id);
  if (elem == null) {
    throw new Error(`Element '${id}' could not be found`);
  }
  return elem;
}
document.addEventListener("DOMContentLoaded", () => {
  const submit = eot("submit");
  const password = eot("password");
  const password_c = eot("password-confirm");
  const message = eot("message");

  submit.addEventListener("click", (e) => {
    if (password.value != password_c.value) {
      e.preventDefault();
      message.innerText = "The password and confirmation aren't matching.";
    }
  });
});
