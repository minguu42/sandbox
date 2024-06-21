INSERT INTO users(id, name)
VALUES ('U_1', 'hsaki');

INSERT INTO repositories(id, owner, name)
VALUES ('REPO_1', 'U_1', 'repo1');

INSERT INTO issues(id, url, title, closed, number, repository)
VALUES ('ISSUE_1', 'http://example.com/repo1/issue/1', 'First Issue', 1, 1, 'REPO_1'),
       ('ISSUE_2', 'http://example.com/repo1/issue/2', 'Second Issue', 0, 2, 'REPO_1'),
       ('ISSUE_3', 'http://example.com/repo1/issue/3', 'Third Issue', 0, 3, 'REPO_1');

INSERT INTO projects(id, title, url, owner)
VALUES ('PJ_1', 'My Project', 'http://example.com/project/1', 'U_1');

INSERT INTO pullrequests(id, base_ref_name, closed, head_ref_name, url, number, repository)
VALUES ('PR_1', 'main', 1, 'feature/kinou1', 'http://example.com/repo1/pr/1', 1, 'REPO_1'),
       ('PR_2', 'main', 0, 'feature/kinou2', 'http://example.com/repo1/pr/2', 2, 'REPO_1');
