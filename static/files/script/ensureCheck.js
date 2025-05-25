function cot(id) {
  const elem = document.getElementById(id);
  if (elem == null) {
    throw new Error(`could not get element with id '${id}'`);
  }
  if (!(elem instanceof HTMLInputElement)) {
    throw new Error(`element with id '${id}' is not an input element`);
  }
  return elem;
}

document.addEventListener("DOMContentLoaded", () => {
  const bc = cot("border-check");
  const pc = cot("papers-check");

  pc.addEventListener("change", () => {
    if (pc.checked) {
        bc.checked = true;
        bc.disabled = true;
    } else {
        bc.disabled = false;
    }
  })
});
