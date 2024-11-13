customElements.define('toast-message', class ToastMessage extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = `
        <style>
            .toast {
                background-color: #5C6BC0;
                color: white;
                border-radius: 0.375rem;
                z-index: 9999;
                opacity: 1;
                padding: 10px 20px;
                margin: 10px;
                font-size: 16px;
                font-weight: 500;
                cursor: pointer;
            }
            .toast.error {
                background-color: red;
            }
            .toast.hide {
                opacity: 0;
                transition: opacity 0.5s ease-in-out, transform 0.5s ease;
                transform: translateY(-50px);
            }
        </style>
        <div class="toast">
            <slot></slot>
        </div>
        `
    }

    connectedCallback() {
        setTimeout(() => {
            this.remove();
        }, 3000)
        this.addEventListener('click', () => this.remove());
    }
    remove() {
        const toast = this.shadowRoot.querySelector('.toast');
        toast.classList.add('hide');
        setTimeout(() => {
            this.parentNode.removeChild(this);
        }, 500);
    }

    setType(type) {
        const toast = this.shadowRoot.querySelector('.toast');
        if (type === 'error') {
            toast.classList.add('error');
        }
    }
});


  
customElements.define('toast-container', class ToastContainer extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: 'open' });
  
      this.shadowRoot.innerHTML = `
        <style>
          :host {
            position: fixed;
            top: 10px;
            right: 10px;
            display: flex;
            flex-direction: column;
            align-items: flex-end;
            pointer-events: auto;
            gap: 10px;
          }
        </style>
        <slot></slot>
      `;
    }
  
    addToast(message, type="") {
      const toast = document.createElement('toast-message');
      toast.textContent = message
      if (type === "error") {
        toast.setType("error");
      }
      this.appendChild(toast);
    }
});