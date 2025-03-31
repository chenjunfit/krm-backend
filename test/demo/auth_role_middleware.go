package demo

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// AuthMiddleware is a middleware for RBAC authorization
type AuthMiddleware struct {
	Warden *ladon.Ladon
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware() *AuthMiddleware {
	warden := &ladon.Ladon{
		Manager: memory.NewMemoryManager(),
	}

	// Setup policies
	setupPolicies(warden)

	return &AuthMiddleware{
		Warden: warden,
	}
}

// setupPolicies initializes all the RBAC policies
func setupPolicies(warden *ladon.Ladon) {
	// Allow all users to create HMD
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "1",
		Description: "All users can create HMDs",
		Subjects:    []string{"<.*>"},
		Resources:   []string{"hmd"},
		Actions:     []string{"create"},
		Effect:      ladon.AllowAccess,
	})

	// Dev role can read all HMDs
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "2",
		Description: "Dev role can read any HMD",
		Subjects:    []string{"role:dev"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"read"},
		Effect:      ladon.AllowAccess,
	})

	// Users can read HMDs they created
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "3",
		Description: "Users can read HMDs they created",
		Subjects:    []string{"<.*>"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"read"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"owner": &ladon.EqualsSubjectCondition{},
		},
	})

	// Users can read HMDs they own
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "4",
		Description: "Users can read HMDs they own",
		Subjects:    []string{"<.*>"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"read"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"is_owner": &IsOwnerCondition{},
		},
	})

	// Admin/operator can read HMDs in same group
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "5",
		Description: "Admin and operator can read HMDs in same group",
		Subjects:    []string{"role:admin", "role:operator"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"read"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"same_group": &SameGroupCondition{},
		},
	})

	// Dev role can update all HMDs
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "6",
		Description: "Dev role can update any HMD",
		Subjects:    []string{"role:dev"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"update"},
		Effect:      ladon.AllowAccess,
	})

	// Admin/operator can update HMDs in same group
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "7",
		Description: "Admin and operator can update HMDs in same group",
		Subjects:    []string{"role:admin", "role:operator"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"update"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"same_group": &SameGroupCondition{},
		},
	})

	// Users can update HMDs they created
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "8",
		Description: "Users can update HMDs they created",
		Subjects:    []string{"<.*>"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"update"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"owner": &ladon.EqualsSubjectCondition{},
		},
	})

	// Users can update HMDs they own
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "9",
		Description: "Users can update HMDs they own",
		Subjects:    []string{"<.*>"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"update"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"is_owner": &IsOwnerCondition{},
		},
	})

	// Regular users can only update Name field
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "10",
		Description: "Regular users can only update Name field",
		Subjects:    []string{"role:user"},
		Resources:   []string{"hmd:<.*>:field"},
		Actions:     []string{"update"},
		Effect:      ladon.DenyAccess,
		Conditions: ladon.Conditions{
			"not_name_field": &NotNameFieldCondition{},
		},
	})

	// Dev role can delete all HMDs
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "11",
		Description: "Dev role can delete any HMD",
		Subjects:    []string{"role:dev"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"delete"},
		Effect:      ladon.AllowAccess,
	})

	// Admin can delete HMDs in same group
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "12",
		Description: "Admin can delete HMDs in same group",
		Subjects:    []string{"role:admin"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"delete"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"same_group": &SameGroupCondition{},
		},
	})

	// Admin can delete HMDs they created
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "13",
		Description: "Admin can delete HMDs they created",
		Subjects:    []string{"role:admin"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"delete"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"owner": &ladon.EqualsSubjectCondition{},
		},
	})

	// Admin can delete HMDs they own
	warden.Manager.Create(context.Background(), &ladon.DefaultPolicy{
		ID:          "14",
		Description: "Admin can delete HMDs they own",
		Subjects:    []string{"role:admin"},
		Resources:   []string{"hmd:<.*>"},
		Actions:     []string{"delete"},
		Effect:      ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"is_owner": &IsOwnerCondition{},
		},
	})
}

// Authorize checks if the user has permission to perform an action
func (m *AuthMiddleware) Authorize(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := getUserFromContext(c)
		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Extract HMD ID from URL if it exists
		hmdID := c.Param("id")
		resource := "hmd"
		if hmdID != "" {
			resource = fmt.Sprintf("hmd:%s", hmdID)
		}

		// Build subject list (user ID + roles)
		subjects := []string{fmt.Sprintf("%d", *user.ID)}
		for _, role := range user.Roles {
			subjects = append(subjects, fmt.Sprintf("role:%s", role.Name))
		}

		// Build context for conditions
		ctx := ladon.Context{
			"owner":      fmt.Sprintf("%d", *user.ID),
			"userID":     *user.ID,
			"userRoles":  getUserRoles(user),
			"userGroups": getUserGroups(user),
		}

		// Add HMD context if applicable
		if hmdID != "" {
			hmd, err := getHMDByID(c, hmdID)
			if err == nil && hmd != nil {
				ctx["hmdCreatorID"] = *hmd.CreatedByID
				ctx["hmdGroups"] = getHMDGroups(hmd)
				ctx["hmdOwners"] = getHMDOwners(hmd)

				// For field-level permissions
				if action == "update" {
					fieldName := c.GetString("fieldName")
					if fieldName != "" {
						ctx["fieldName"] = fieldName
						resource = fmt.Sprintf("%s:field", resource)
					}
				}
			}
		}

		// Check permissions for each subject
		var allowed bool
		for _, subject := range subjects {
			err := m.Warden.IsAllowed(&ladon.Request{
				Subject:  subject,
				Resource: resource,
				Action:   action,
				Context:  ctx,
			})

			if err == nil {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}

		c.Next()
	}
}

// Custom conditions for Ladon
// IsOwnerCondition checks if the user is in the HMD owners list
type IsOwnerCondition struct{}

func (c *IsOwnerCondition) GetName() string {
	return "IsOwnerCondition"
}

func (c *IsOwnerCondition) Fulfills(value interface{}, r *ladon.Request) bool {
	userID := r.Context["userID"].(uint)
	owners := r.Context["hmdOwners"].([]uint)

	for _, ownerID := range owners {
		if userID == ownerID {
			return true
		}
	}

	return false
}

// SameGroupCondition checks if the user and HMD share at least one group
type SameGroupCondition struct{}

func (c *SameGroupCondition) GetName() string {
	return "SameGroupCondition"
}

func (c *SameGroupCondition) Fulfills(value interface{}, r *ladon.Request) bool {
	userGroups := r.Context["userGroups"].([]uint)
	hmdGroups := r.Context["hmdGroups"].([]uint)

	for _, userGroup := range userGroups {
		for _, hmdGroup := range hmdGroups {
			if userGroup == hmdGroup {
				return true
			}
		}
	}

	return false
}

// NotNameFieldCondition checks if the field being updated is not "Name"
type NotNameFieldCondition struct{}

func (c *NotNameFieldCondition) GetName() string {
	return "NotNameFieldCondition"
}

func (c *NotNameFieldCondition) Fulfills(value interface{}, r *ladon.Request) bool {
	fieldName, ok := r.Context["fieldName"].(string)
	if !ok {
		return true // If no field specified, deny
	}

	return fieldName != "Name"
}

// Helper functions
func getUserFromContext(c *gin.Context) *User {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil
	}
	user, ok := userInterface.(*User)
	if !ok {
		return nil
	}
	return user
}

func getHMDByID(c *gin.Context, idStr string) (*HMD, error) {
	// This would typically use your database access layer
	// Mock implementation for example purposes
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	// You would fetch from DB here
	var hmd HMD
	if result := c.MustGet("db").(*gorm.DB).First(&hmd, id); result.Error != nil {
		return nil, result.Error
	}

	return &hmd, nil
}

func getUserRoles(user *User) []string {
	roles := make([]string, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}
	return roles
}

func getUserGroups(user *User) []uint {
	groups := make([]uint, 0, len(user.Groups))
	for _, group := range user.Groups {
		groups = append(groups, *group.ID)
	}
	return groups
}

func getHMDGroups(hmd *HMD) []uint {
	groups := make([]uint, 0, len(hmd.Groups))
	for _, group := range hmd.Groups {
		groups = append(groups, *group.ID)
	}
	return groups
}

func getHMDOwners(hmd *HMD) []uint {
	owners := make([]uint, 0, len(hmd.Owners))
	for _, owner := range hmd.Owners {
		owners = append(owners, *owner.ID)
	}
	return owners
}
