package report

import (
	"context"
	"strings"

	"github.com/olekukonko/tablewriter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"os"

	managementv3 "github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/sirupsen/logrus"
)

// global variable to store userDetails to speed up the userlookup at report generation //
var userDetailsMap map[string]string

type ReportCommandConfig struct {
	Cluster      string
	Client       managementv3.Interface
	Context      context.Context
	ReportFormat string
}

type Report struct {
	GlobalRoles []managementv3.GlobalRoleBinding
	Clusters    map[string]ClusterDetails
}

type ClusterDetails struct {
	ClusterName  string
	ClusterRoles []managementv3.ClusterRoleTemplateBinding
	Projects     map[string]ProjectDetails
}

type ProjectDetails struct {
	ProjectName  string
	ProjectRoles []managementv3.ProjectRoleTemplateBinding
}

//GenerateReport will generate the rbac report for the cluster
func (rc ReportCommandConfig) GenerateReport() {

	userDetailsMap = make(map[string]string)
	userDetailsMap = rc.getUserList()

	report := Report{}
	clusterDetailsMap := make(map[string]ClusterDetails)
	clusterList, err := rc.listClusters()
	if err != nil {
		logrus.Errorf("Error while fetching cluster list %v\n", err)
		os.Exit(1)
	}

	for cluster, clusterName := range clusterList {
		clusterDetails := ClusterDetails{}
		projectDetails := ProjectDetails{}
		projectRoleMap := make(map[string]ProjectDetails)
		projectList, err := rc.listProjects(cluster)
		if err != nil {
			logrus.Errorf("Error ranging over projects for %s: %v\n", cluster, err)
			os.Exit(1)
		}
		for _, project := range projectList {
			prtblist, err := rc.listProjectMembers(project.Name)
			if err != nil {
				logrus.Errorf("Error while fetching project members for %s in %s %v\n", project, cluster, err)
				os.Exit(1)
			}
			projectDetails.ProjectName = project.Spec.DisplayName
			projectDetails.ProjectRoles = prtblist
			projectRoleMap[project.Name] = projectDetails
		}

		clusterRolesList, err := rc.listClusterMembers(cluster)
		if err != nil {
			logrus.Errorf("Error fetching clusterroles for cluster %s: %v\n", cluster, err)
			os.Exit(1)
		}
		clusterDetails.ClusterRoles = clusterRolesList
		clusterDetails.Projects = projectRoleMap
		clusterDetails.ClusterName = clusterName
		clusterDetailsMap[cluster] = clusterDetails
	}

	grtb, err := rc.getGlobalRoleBindings()
	if err != nil {
		logrus.Errorf("Error fetching global role bindings %v\n", err)
		os.Exit(1)
	}
	report.GlobalRoles = grtb
	report.Clusters = clusterDetailsMap
	rc.generateTable(report)
}

func (rc ReportCommandConfig) listClusters() (clusterDetails map[string]string, err error) {
	clusterDetails = make(map[string]string)
	cl, err := rc.Client.Clusters("").List(metav1.ListOptions{})
	if err != nil {
		return
	}

	for _, clusterName := range cl.Items {
		clusterDetails[clusterName.Name] = clusterName.Spec.DisplayName
	}

	return
}

func (rc ReportCommandConfig) listProjects(cluster string) (projectList []managementv3.Project, err error) {
	plist, err := rc.Client.Projects(cluster).List(metav1.ListOptions{})
	if err != nil {
		return
	}

	projectList = plist.Items
	return
}

func (rc ReportCommandConfig) listProjectMembers(project string) (projectMembers []managementv3.ProjectRoleTemplateBinding, err error) {
	prtblist, err := rc.Client.ProjectRoleTemplateBindings(project).List(metav1.ListOptions{})
	if err != nil {
		return
	}
	projectMembers = prtblist.Items

	return
}

func (rc ReportCommandConfig) listClusterMembers(cluster string) (clusterMembers []managementv3.ClusterRoleTemplateBinding, err error) {
	crtblist, err := rc.Client.ClusterRoleTemplateBindings(cluster).List(metav1.ListOptions{})
	if err != nil {
		return
	}

	clusterMembers = crtblist.Items

	return
}

func (rc ReportCommandConfig) getRoleTemplateName(roleTemplate string) (roleTemplateName string, err error) {
	rt, err := rc.Client.RoleTemplates("").Get(roleTemplate, metav1.GetOptions{})
	if err != nil {
		return
	}
	roleTemplateName = rt.DisplayName
	return
}

func (rc ReportCommandConfig) getUserList() (userDetails map[string]string) {
	userDetails = make(map[string]string)
	userList, err := rc.Client.Users("").List(metav1.ListOptions{})
	if err != nil {
		return
	}
	for _, user := range userList.Items {
		userDetails[user.ObjectMeta.Name] = user.DisplayName
	}
	return
}

func (rc ReportCommandConfig) getGlobalRoleBindings() (globalRoles []managementv3.GlobalRoleBinding, err error) {
	grtb, err := rc.Client.GlobalRoleBindings("").List(metav1.ListOptions{})
	if err != nil {
		return
	}
	globalRoles = grtb.Items
	return
}

func fetchUserDisplayName(userName string) (displayName string) {
	if val, ok := userDetailsMap[userName]; ok {
		displayName = val
	} else {
		displayName = userName
	}

	return
}

// generateTable will render the table for the cluster report
func (rc ReportCommandConfig) generateTable(r Report) {

	generateGlobalRolesTable(r)
	if len(rc.Cluster) != 0 {
		for _, cd := range r.Clusters {
			if cd.ClusterName == rc.Cluster {
				rc.parseClusters(cd)
			}
		}
	} else {
		for _, cd := range r.Clusters {
			rc.parseClusters(cd)
		}
	}
}

func (rc ReportCommandConfig) parseClusters(cd ClusterDetails) {
	clusterRolesTable := tablewriter.NewWriter(os.Stdout)
	projectRolesTable := tablewriter.NewWriter(os.Stdout)
	clusterName := cd.ClusterName
	clusterRolesTable.SetHeader([]string{"Cluster", "User", "Group", "Group Principal ID", "Cluster Role"})
	projectRolesTable.SetHeader([]string{"Cluster", "Project", "User", "Group", "Group Principal ID", "Project Role"})

	//Add cluster table rows //
	for _, clusterRole := range cd.ClusterRoles {
		var roleTemplateName string
		var err error
		if strings.Contains(clusterRole.RoleTemplateName, "rt-") {
			roleTemplateName, err = rc.getRoleTemplateName(clusterRole.RoleTemplateName)
			if err != nil {
				roleTemplateName = clusterRole.RoleTemplateName
			}
		} else {
			roleTemplateName = clusterRole.RoleTemplateName
		}
		clusterRolesTable.Append([]string{clusterName, fetchUserDisplayName(clusterRole.UserName),
			clusterRole.GroupName, clusterRole.GroupPrincipalName, roleTemplateName})
	}

	// Add project table rows //
	for _, projectRolesList := range cd.Projects {
		project := projectRolesList.ProjectName
		for _, projectRoles := range projectRolesList.ProjectRoles {
			var roleTemplateName string
			var err error
			if strings.Contains(projectRoles.RoleTemplateName, "rt-") {
				roleTemplateName, err = rc.getRoleTemplateName(projectRoles.RoleTemplateName)
				if err != nil {
					roleTemplateName = projectRoles.RoleTemplateName
				}
			} else {
				roleTemplateName = projectRoles.RoleTemplateName
			}
			projectRolesTable.Append([]string{clusterName, project, fetchUserDisplayName(projectRoles.UserName),
				projectRoles.GroupName, projectRoles.GroupPrincipalName,
				roleTemplateName})
		}
	}
	clusterRolesTable.Render()
	projectRolesTable.Render()
}

func generateGlobalRolesTable(r Report) {
	globalRolesTable := tablewriter.NewWriter(os.Stdout)
	globalRolesTable.SetHeader([]string{"User", "Global Role"})
	for _, grtb := range r.GlobalRoles {
		globalRolesTable.Append([]string{fetchUserDisplayName(grtb.UserName), grtb.GlobalRoleName})
	}
	globalRolesTable.Render()
}
