var socket = new Ws("ws://localhost:8080/readingList/sync");

socket.On("saved", function () {
  console.log("receive: on saved");
  fetchbooks(function (items) {
    app.books = items
  });
});


function fetchbooks(onComplete) {
  axios.get("/readingList").then(response => {
    if (response.data === null) {
      return;
    }
    console.log("we are here")
    onComplete(response.data);
  });
}

var bookStorage = {
  fetch: function () {
    var books = [];
    fetchbooks(function (items) {
      for (var i = 0; i < items.length; i++) {
        books.push(items[i]);
      }
    });
    return books;
  },
  save: function (books) {
    axios.post("/readingList", JSON.stringify(books)).then(response => {
      if (!response.data.success) {
        window.alert("saving had a failure");
        return;
      }
      // console.log("send: save");
      socket.Emit("save")
    });
  }
}

// filters
var filters = {
  all: function (books) {
    return books
  },
  active: function (books) {
    return books.filter(function (book) {
      return !book.completed
    })
  },
  completed: function (books) {
    return books.filter(function (book) {
      return book.completed
    })
  }
}

// app Vue instance
var app = new Vue({
  // app initial state
  data: {
    books: bookStorage.fetch(),
    newBook: '',
    editedBook: null,
    visibility: 'all'
  },

  computed: {
    filteredBooks: function () {
      return filters[this.visibility](this.books)
    },
    remaining: function () {
      return filters.active(this.books).length
    },
    allDone: {
      get: function () {
        return this.remaining === 0
      },
      set: function (value) {
        this.books.forEach(function (book) {
          book.completed = value
        })
        this.notifyChange();
      }
    }
  },

  filters: {
    pluralize: function (n) {
      return n === 1 ? 'book' : 'books'
    }
  },


  methods: {
    notifyChange: function () {
      bookStorage.save(this.books)
    },
    addBook: function () {
      var value = this.newBook && this.newBook.trim()
      if (!value) {
        return
      }
      this.books.push({
        id: this.books.length + 1, // just for the client-side.
        title: value,
        completed: false
      })
      this.newBook = ''
      this.notifyChange();
    },

    completeBook: function (book) {
      if (book.completed) {
        book.completed = false;
      } else {
        book.completed = true;
      }
      this.notifyChange();
    },
    removeBook: function (book) {
      this.books.splice(this.books.indexOf(book), 1)
      this.notifyChange();
    },

    editBook: function (book) {
      this.beforeEditTitle = book.title
      this.beforeEditAuthor = book.author
      this.editedBook = book
    },

    doneEdit: function (book) {
      if (!this.editedBook) {
        return
      }
      this.editedBook = null
      book.title = book.title.trim();
      if (!book.title) {
        this.removeBook(book);
      }
      this.notifyChange();
    },

    cancelEdit: function (book) {
      this.editedBook = null
      book.title = this.beforeEditTitle
      book.author = this.beforeEditAuthor
    },

    removeCompleted: function () {
      this.books = filters.active(this.books);
      this.notifyChange();
    }
  },

  directives: {
    'book-focus': function (el, binding) {
      if (binding.value) {
        el.focus()
      }
    }
  }
})

// handle routing
function onHashChange() {
  var visibility = window.location.hash.replace(/#\/?/, '')
  if (filters[visibility]) {
    app.visibility = visibility
  } else {
    window.location.hash = ''
    app.visibility = 'all'
  }
}

window.addEventListener('hashchange', onHashChange)
onHashChange()

// mount
app.$mount('.readinglistapp')