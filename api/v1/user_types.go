/*
Copyright 2024 382023823@qq.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UserSpec defines the desired state of User
type UserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of User. Edit user_types.go to remove/update
	CreateBy   string      `json:"createBy,omitempty"`
	CreateTime metav1.Time `json:"createTime,omitempty"`
	UpdateBy   string      `json:"updateBy,omitempty"`
	UpdateTime metav1.Time `json:"updateTime,omitempty"`
	Status     string      `json:"status,omitempty"`
	UserName   string      `json:"userName,omitempty"`
	UserGender string      `json:"userGender,omitempty"`
	NickName   string      `json:"nickName,omitempty"`
	UserPhone  string      `json:"userPhone,omitempty"`
	UserEmail  string      `json:"userEmail,omitempty"`
	UserRoles  []string    `json:"userRoles,omitempty"`
}

// UserStatus defines the observed state of User
type UserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status bool `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type="boolean",JSONPath=".status.status",description="keycloak uerid"
// +kubebuilder:printcolumn:name="用户名",type="string",JSONPath=".spec.userName",description="username"
// +kubebuilder:printcolumn:name="性别",type="string",JSONPath=".spec.userGender",description="gender"
// +kubebuilder:printcolumn:name="小名",type="string",JSONPath=".spec.nickName",description="nickname"
// +kubebuilder:printcolumn:name="电话",type="string",JSONPath=".spec.userPhone",description="cellphone"
// +kubebuilder:printcolumn:name="创建人",type="string",JSONPath=".spec.createBy",description="createby"
// +kubebuilder:printcolumn:name="邮箱",type="string",JSONPath=".spec.userEmail",description="Email"
// +kubebuilder:printcolumn:name="创建时间",type="date",JSONPath=".metadata.creationTimestamp",description="创建时间"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// User is the Schema for the users API
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of User
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
