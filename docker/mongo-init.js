// MongoDB initialization script
db = db.getSiblingDB('project_planton');

// Create collections if they don't exist
db.createCollection('deployment_components');

// Create indexes for better performance
db.deployment_components.createIndex({ "provider": 1 });
db.deployment_components.createIndex({ "kind": 1 });
db.deployment_components.createIndex({ "name": 1 });
db.deployment_components.createIndex({ "createdAt": 1 });

print('MongoDB initialized successfully for project_planton database');
