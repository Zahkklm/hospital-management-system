class DashboardManager {
    constructor() {
        this.patients = [];
        this.currentUser = null;
        this.isEditing = false;
        
        this.initializeAuth();
        this.initializeEventListeners();
        this.loadUserInfo();
        this.loadPatients();
    }

    initializeAuth() {
        const token = localStorage.getItem('token');
        if (!token) {
            window.location.href = '/login';
            return;
        }

        const userData = localStorage.getItem('user');
        if (userData) {
            this.currentUser = JSON.parse(userData);
        }
    }

    initializeEventListeners() {
        // Patient form submission
        const patientForm = document.getElementById('patientForm');
        if (patientForm) {
            patientForm.addEventListener('submit', (e) => this.handlePatientSubmit(e));
        }

        // Modal close on background click
        const modal = document.getElementById('patientModal');
        if (modal) {
            modal.addEventListener('click', (e) => {
                if (e.target === modal) {
                    this.closePatientModal();
                }
            });
        }

        // Keyboard shortcuts
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape') {
                this.closePatientModal();
            }
        });
    }

    loadUserInfo() {
        if (this.currentUser) {
            document.getElementById('userName').textContent = this.currentUser.username;
            document.getElementById('userRole').textContent = this.currentUser.role.toUpperCase();
            document.getElementById('userAvatar').textContent = this.currentUser.username.charAt(0).toUpperCase();
            
            // Hide add button for doctors if needed
            if (this.currentUser.role === 'doctor') {
                const addBtn = document.getElementById('addPatientBtn');
                if (addBtn) {
                    addBtn.style.display = 'none';
                }
            }
        }
    }

    async loadPatients() {
        const loading = document.getElementById('loading');
        const patientList = document.getElementById('patient-list');
        const noPatients = document.getElementById('no-patients');

        loading.classList.remove('d-none');
        patientList.innerHTML = '';
        noPatients.classList.add('d-none');

        try {
            const token = localStorage.getItem('token');
            const response = await fetch('/api/patients', {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (!response.ok) {
                throw new Error('Failed to fetch patients');
            }

            const patients = await response.json();
            this.patients = patients || [];

            if (this.patients.length === 0) {
                noPatients.classList.remove('d-none');
            } else {
                this.renderPatients();
            }
        } catch (error) {
            console.error('Error loading patients:', error);
            this.showAlert('Failed to load patients. Please try again.', 'danger');
        } finally {
            loading.classList.add('d-none');
        }
    }

    renderPatients() {
        const patientList = document.getElementById('patient-list');
        patientList.innerHTML = '';

        this.patients.forEach(patient => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${patient.id}</td>
                <td>${this.escapeHtml(patient.first_name)} ${this.escapeHtml(patient.last_name)}</td>
                <td>${this.formatDate(patient.date_of_birth)}</td>
                <td>${this.calculateAge(patient.date_of_birth)}</td>
                <td>${this.capitalize(patient.gender)}</td>
                <td>${this.escapeHtml(patient.phone || 'N/A')}</td>
                <td>${this.escapeHtml(patient.email || 'N/A')}</td>
                <td>
                    <div class="table-actions">
                        <button class="btn btn-primary btn-sm" onclick="dashboard.editPatient(${patient.id})">
                            ✏️ Edit
                        </button>
                        ${this.currentUser?.role === 'receptionist' ? `
                            <button class="btn btn-danger btn-sm" onclick="dashboard.deletePatient(${patient.id})">
                                🗑️ Delete
                            </button>
                        ` : ''}
                    </div>
                </td>
            `;
            patientList.appendChild(row);
        });
    }

    showAddPatientModal() {
        this.isEditing = false;
        document.getElementById('modalTitle').textContent = 'Add New Patient';
        document.getElementById('saveText').textContent = 'Add Patient';
        document.getElementById('patientForm').reset();
        document.getElementById('patientId').value = '';
        document.getElementById('patientModal').classList.add('show');
    }

    async editPatient(id) {
        const patient = this.patients.find(p => p.id === id);
        if (!patient) {
            this.showAlert('Patient not found', 'danger');
            return;
        }

        this.isEditing = true;
        document.getElementById('modalTitle').textContent = 'Edit Patient';
        document.getElementById('saveText').textContent = 'Update Patient';
        
        // Populate form
        document.getElementById('patientId').value = patient.id;
        document.getElementById('firstName').value = patient.first_name;
        document.getElementById('lastName').value = patient.last_name;
        document.getElementById('dateOfBirth').value = patient.date_of_birth;
        document.getElementById('gender').value = patient.gender;
        document.getElementById('phone').value = patient.phone || '';
        document.getElementById('email').value = patient.email || '';
        document.getElementById('address').value = patient.address || '';
        
        document.getElementById('patientModal').classList.add('show');
    }

    async deletePatient(id) {
        if (!confirm('Are you sure you want to delete this patient? This action cannot be undone.')) {
            return;
        }

        try {
            const token = localStorage.getItem('token');
            const response = await fetch(`/api/patients/${id}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.ok) {
                this.showAlert('Patient deleted successfully', 'success');
                this.loadPatients();
            } else {
                throw new Error('Failed to delete patient');
            }
        } catch (error) {
            console.error('Error deleting patient:', error);
            this.showAlert('Failed to delete patient. Please try again.', 'danger');
        }
    }

    async handlePatientSubmit(event) {
        event.preventDefault();
        
        const saveBtn = document.getElementById('savePatientBtn');
        const saveText = document.getElementById('saveText');
        const saveSpinner = document.getElementById('saveSpinner');
        
        this.setLoadingState(saveBtn, saveText, saveSpinner, true);

        const formData = {
            first_name: document.getElementById('firstName').value.trim(),
            last_name: document.getElementById('lastName').value.trim(),
            date_of_birth: document.getElementById('dateOfBirth').value,
            gender: document.getElementById('gender').value,
            phone: document.getElementById('phone').value.trim(),
            email: document.getElementById('email').value.trim(),
            address: document.getElementById('address').value.trim()
        };

        // Validation
        if (!this.validatePatientForm(formData)) {
            this.setLoadingState(saveBtn, saveText, saveSpinner, false);
            return;
        }

        try {
            const token = localStorage.getItem('token');
            const patientId = document.getElementById('patientId').value;
            
            const url = this.isEditing ? `/api/patients/${patientId}` : '/api/patients';
            const method = this.isEditing ? 'PUT' : 'POST';

            const response = await fetch(url, {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(formData)
            });

            if (response.ok) {
                const message = this.isEditing ? 'Patient updated successfully' : 'Patient added successfully';
                this.showAlert(message, 'success');
                this.closePatientModal();
                this.loadPatients();
            } else {
                const error = await response.json();
                throw new Error(error.error || 'Failed to save patient');
            }
        } catch (error) {
            console.error('Error saving patient:', error);
            this.showAlert(error.message || 'Failed to save patient. Please try again.', 'danger');
        } finally {
            this.setLoadingState(saveBtn, saveText, saveSpinner, false);
        }
    }

    validatePatientForm(formData) {
        if (!formData.first_name || !formData.last_name || !formData.date_of_birth || !formData.gender) {
            this.showAlert('Please fill in all required fields', 'danger');
            return false;
        }

        // Validate date of birth
        const dob = new Date(formData.date_of_birth);
        const today = new Date();
        
        if (dob > today) {
            this.showAlert('Date of birth cannot be in the future', 'danger');
            return false;
        }

        const age = this.calculateAge(formData.date_of_birth);
        if (age > 150) {
            this.showAlert('Please enter a valid date of birth', 'danger');
            return false;
        }

        // Validate email if provided
        if (formData.email && !this.isValidEmail(formData.email)) {
            this.showAlert('Please enter a valid email address', 'danger');
            return false;
        }

        return true;
    }

    closePatientModal() {
        document.getElementById('patientModal').classList.remove('show');
        document.getElementById('patientForm').reset();
        this.isEditing = false;
    }

    // Utility functions
    calculateAge(dateOfBirth) {
        const today = new Date();
        const birthDate = new Date(dateOfBirth);
        let age = today.getFullYear() - birthDate.getFullYear();
        const monthDiff = today.getMonth() - birthDate.getMonth();
        
        if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
            age--;
        }
        
        return age;
    }

    formatDate(dateString) {
        const date = new Date(dateString);
        return date.toLocaleDateString('en-US', {
            year: 'numeric',
            month: 'short',
            day: 'numeric'
        });
    }

    capitalize(str) {
        return str ? str.charAt(0).toUpperCase() + str.slice(1) : '';
    }

    escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    isValidEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
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
        alert.innerHTML = `
            <div style="display: flex; justify-content: space-between; align-items: center;">
                <span>${message}</span>
                <button onclick="this.parentElement.parentElement.remove()" style="background: none; border: none; font-size: 1.2rem; cursor: pointer;">&times;</button>
            </div>
        `;

        alertContainer.appendChild(alert);

        // Auto-hide success/info alerts
        if (type === 'success' || type === 'info') {
            setTimeout(() => {
                alert.remove();
            }, 5000);
        }
    }
}

// Global functions for onclick handlers
window.showAddPatientModal = () => dashboard.showAddPatientModal();
window.closePatientModal = () => dashboard.closePatientModal();

// Logout function
window.logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = '/login';
};

// Initialize dashboard when DOM is loaded
let dashboard;
document.addEventListener('DOMContentLoaded', () => {
    dashboard = new DashboardManager();
});