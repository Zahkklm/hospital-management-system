class AuthManager {
    constructor() {
        this.initializeEventListeners();
        this.checkExistingAuth();
    }

    initializeEventListeners() {
        // Login form
        const loginForm = document.getElementById('loginForm');
        if (loginForm) {
            loginForm.addEventListener('submit', (e) => this.handleLogin(e));
        }

        // Register form
        const registerForm = document.getElementById('registerForm');
        if (registerForm) {
            registerForm.addEventListener('submit', (e) => this.handleRegister(e));
            
            // Password confirmation validation
            const confirmPassword = document.getElementById('confirmPassword');
            if (confirmPassword) {
                confirmPassword.addEventListener('input', () => this.validatePasswordMatch());
            }
        }
    }

    checkExistingAuth() {
        const token = localStorage.getItem('token');
        if (token && (window.location.pathname === '/login' || window.location.pathname === '/register')) {
            window.location.href = '/api/dashboard';
        }
    }

    async handleLogin(event) {
        event.preventDefault();
        
        const loginBtn = document.getElementById('loginBtn');
        const loginText = document.getElementById('loginText');
        const loginSpinner = document.getElementById('loginSpinner');
        
        // Show loading state
        this.setLoadingState(loginBtn, loginText, loginSpinner, true);
        
        const formData = {
            username: document.getElementById('username').value.trim(),
            password: document.getElementById('password').value
        };

        // Basic validation
        if (!formData.username || !formData.password) {
            this.showAlert('Please fill in all fields', 'danger');
            this.setLoadingState(loginBtn, loginText, loginSpinner, false);
            return;
        }

        try {
            const response = await fetch('/api/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData)
            });

            const data = await response.json();

            if (data.success && data.token) {
                localStorage.setItem('token', data.token);
                localStorage.setItem('user', JSON.stringify(data.user));
                
                this.showAlert('Login successful! Redirecting...', 'success');
                
                setTimeout(() => {
                    window.location.href = '/api/dashboard';
                }, 1000);
            } else {
                this.showAlert(data.error || 'Login failed. Please check your credentials.', 'danger');
            }
        } catch (error) {
            console.error('Login error:', error);
            this.showAlert('Connection error. Please try again.', 'danger');
        } finally {
            this.setLoadingState(loginBtn, loginText, loginSpinner, false);
        }
    }

    async handleRegister(event) {
        event.preventDefault();
        
        const registerBtn = document.getElementById('registerBtn');
        const registerText = document.getElementById('registerText');
        const registerSpinner = document.getElementById('registerSpinner');
        
        // Show loading state
        this.setLoadingState(registerBtn, registerText, registerSpinner, true);
        
        const formData = {
            username: document.getElementById('username').value.trim(),
            password: document.getElementById('password').value,
            role: document.getElementById('role').value
        };

        const confirmPassword = document.getElementById('confirmPassword').value;

        // Validation
        if (!this.validateRegistrationForm(formData, confirmPassword)) {
            this.setLoadingState(registerBtn, registerText, registerSpinner, false);
            return;
        }

        try {
            const response = await fetch('/api/auth/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData)
            });

            const data = await response.json();

            if (data.success) {
                this.showAlert('Account created successfully! Redirecting to login...', 'success');
                
                setTimeout(() => {
                    window.location.href = '/login';
                }, 2000);
            } else {
                this.showAlert(data.error || 'Registration failed. Please try again.', 'danger');
            }
        } catch (error) {
            console.error('Registration error:', error);
            this.showAlert('Connection error. Please try again.', 'danger');
        } finally {
            this.setLoadingState(registerBtn, registerText, registerSpinner, false);
        }
    }

    validateRegistrationForm(formData, confirmPassword) {
        // Check all fields are filled
        if (!formData.username || !formData.password || !formData.role || !confirmPassword) {
            this.showAlert('Please fill in all fields', 'danger');
            return false;
        }

        // Username validation
        if (formData.username.length < 3) {
            this.showAlert('Username must be at least 3 characters long', 'danger');
            return false;
        }

        // Password validation
        if (!this.validatePassword(formData.password)) {
            return false;
        }

        // Password confirmation
        if (formData.password !== confirmPassword) {
            this.showAlert('Passwords do not match', 'danger');
            return false;
        }

        return true;
    }

    validatePassword(password) {
        if (password.length < 8) {
            this.showAlert('Password must be at least 8 characters long', 'danger');
            return false;
        }

        const hasUpper = /[A-Z]/.test(password);
        const hasLower = /[a-z]/.test(password);
        const hasNumber = /\d/.test(password);
        const hasSpecial = /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password);

        if (!hasUpper || !hasLower || !hasNumber || !hasSpecial) {
            this.showAlert('Password must contain uppercase, lowercase, number and special character', 'danger');
            return false;
        }

        return true;
    }

    validatePasswordMatch() {
        const password = document.getElementById('password').value;
        const confirmPassword = document.getElementById('confirmPassword').value;
        const confirmField = document.getElementById('confirmPassword');

        if (confirmPassword && password !== confirmPassword) {
            confirmField.style.borderColor = 'var(--danger-color)';
        } else {
            confirmField.style.borderColor = 'var(--border-color)';
        }
    }

    setLoadingState(button, textElement, spinnerElement, isLoading) {
        if (isLoading) {
            button.disabled = true;
            textElement.classList.add('d-none');
            spinnerElement.classList.remove('d-none');
        } else {
            button.disabled = false;
            textElement.classList.remove('d-none');
            spinnerElement.classList.add('d-none');
        }
    }

    showAlert(message, type = 'info') {
        const alertContainer = document.getElementById('alert-container');
        if (!alertContainer) return;

        const alert = document.createElement('div');
        alert.className = `alert alert-${type}`;
        alert.textContent = message;

        alertContainer.innerHTML = '';
        alertContainer.appendChild(alert);

        // Auto-hide success/info alerts
        if (type === 'success' || type === 'info') {
            setTimeout(() => {
                alert.remove();
            }, 5000);
        }
    }
}

// Initialize auth manager when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new AuthManager();
});