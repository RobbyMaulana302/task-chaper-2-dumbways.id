// data contact.html
async function submitData(event) {
  event.preventDefault();

  let name = await document.querySelector("#input-name").value;
  let email = await document.querySelector("#input-email").value;
  let phoneNumber = await document.querySelector("#input-phone-number").value;
  let subject = await document.querySelector("#input-subject").value;
  let message = await document.querySelector("#input-message").value;

  if (name === "") {
    alert("Required!! Name can not null");
  } else if (email === "") {
    alert("Required!! Email can not null");
  } else if (phoneNumber === "") {
    alert("Requred!! Phone number can not null");
  } else if (message === "") {
    alert("Required!! Message can not null");
  }

  let emailReceiver = "ccngkremi1@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Hallo, nama saya ${name}, ${message}. silakan kontak saya ke nomer ${phoneNumber}`;
  a.click();

  let emailer = {
    name,
    email,
    phoneNumber,
    subject,
    message,
  };

  console.log(emailer);
}

//data addMyProject.html
let dataProject = [];

function addProject() {
  let title = document.querySelector("#input-project-name").value;
  let startDate = document.querySelector("#input-start-date").value;
  let endDate = document.querySelector("#input-end-date").value;
  let description = document.querySelector("#input-description").value;
  let tech = document.getElementsByName("checkbox-technology");
  let image = document.querySelector("#input-image").files;

  let technology = [];
  for (let index = 0; index < tech.length; index++) {
    if (tech[index].checked) {
      technology.push(tech[index].value);
    }
  }

  image = URL.createObjectURL(image[0]);

  let project = {
    title,
    startDate,
    endDate,
    description,
    technology,
    image,
  };

  dataProject.push(project);
  renderProject();
}

function renderProject() {
  document.getElementById("my-project").innerHTML = "";

  for (let index = 0; index < dataProject.length; index++) {
    let tech = "";

    for (let k = 0; k < dataProject[index].technology.length; k++) {
      tech += `
      <img
          src="assets/images/${dataProject[index].technology[k]}.png"
          alt=""
          class="rounded-circle img-thumbnail"
          />
      <img
    `;
    }

    document.getElementById("my-project").innerHTML += `
      <div class="col-md-4 mb-5">
          <div class="card p-2">
            <img
              src="assets/images/${dataProject[index].image}.png"
              alt="project-image"
              class="img-thumbnail"
            />
            <div class="card-body">
              <h5 class="card-title fw-bold">${dataProject[index].title}</h5>
              <p class="text-muted card-text">${durationTime(
                dataProject[index].startDate,
                dataProject[index].endDate
              )}</p>
              <p>
                ${dataProject[index].description}
              </p>
              ${tech}
            </div>
            <div class="row">
              <div class="col">
                <a href="#" class="btn btn-dark w-100">edit</a>
              </div>
              <div class="col">
                <a href="#" class="btn btn-dark w-100">delete</a>
              </div>
            </div>
          </div>
       </div>
  `;
  }
  //   dataProject.forEach((element) => {});
}

function durationTime(startDate, endDate) {
  let start = new Date(startDate);
  let end = new Date(endDate);
  let distance = start - end;
  let milisecond = 1000;
  let secondInHour = 3600;
  let hourInDay = 24;
  let dayInWeek = 7;
  let dayInMonth = 30;
  let dayInYear = 365;

  let distanceYear = Math.floor(
    distance / (milisecond * secondInHour * hourInDay * dayInYear)
  );

  let distanceMonth = Math.floor(
    distance / (milisecond * secondInHour * hourInDay * dayInMonth)
  );
  let distanceWeek = Math.floor(
    distance / (milisecond * secondInHour * hourInDay * dayInWeek)
  );
  let distanceDay = Math.floor(
    distance / (milisecond * secondInHour * hourInDay)
  );

  if (distanceDay < 7) {
    return `${distanceDay} days ago`;
  } else if (distanceDay < 30) {
    return `${distanceWeek} weeks ago`;
  } else if (distanceDay < 365) {
    return `${distanceMonth} months ago`;
  } else {
    `${distanceYear} years ago`;
  }
}

// data testimonial.html
const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://api.npoint.io/b89342520bde88ffb2c9");
  xhr.onload = () => {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response));
    } else {
      reject("Error loading data!");
    }
  };
  xhr.onerror = () => {
    reject("Network Error");
  };
  xhr.send();
});

async function getData() {
  const response = await promise;
  let testimonialHTML = "";

  response.forEach((element) => {
    testimonialHTML += `
    <div class="col-md-4 mb-4">
        <div class="card p-2 shadow-sm d-flex">
        <img
            src="${element.image}"
        />
        <p class="fst-italic text-muted mt-3">"${element.quote}"</p>
        <p class="fw-bold text-end px-3">- ${element.name}</p>
        <p class="text-end px-3">${element.rating} <i class="fa-solid fa-star"></i></p>
        </div>
    </div>
    `;
  });

  document.querySelector(".testimonial").innerHTML = testimonialHTML;
}

getData();

async function getRating(rating) {
  const response = await promise;
  let testimonialHTML = "";

  const dataFilter = response.filter((data) => {
    return data.rating === rating;
  });

  if (dataFilter.length === 0) {
    testimonialHTML += "<h1>Data not found!</h1>";
  } else {
    dataFilter.forEach((element) => {
      testimonialHTML += `
      <div class="col-md-4 mb-4">
        <div class="card p-2 shadow-sm d-flex">
        <img
            src="${element.image}"
        />
        <p class="fst-italic text-muted mt-3">"${element.quote}"</p>
        <p class="fw-bold text-end px-3">- ${element.name}</p>
        <p class="text-end px-3">${element.rating} <i class="fa-solid fa-star"></i></p>
        </div>
    </div>
      `;
    });
  }

  document.querySelector(".testimonial").innerHTML = testimonialHTML;
}

function alertDelete() {
  let ok = document.querySelector("#delet");
  ok.addEventListener("click", () => {
    return confirm("bro keluar kek");
  });
}
